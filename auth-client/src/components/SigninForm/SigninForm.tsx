import { Button, Form, FormInput } from "@demo.io/utils/dist/components"
import { Error } from "@demo.io/utils/dist/types"
import {
  addToLocalstorage,
  getFromLocalStorage,
} from "@demo.io/utils/dist/utils"
import Link from "next/link"
import { FC, useEffect, useState } from "react"
import { signinRequest } from "./signinRequest"

const SinginForm: FC = () => {
  const [errors, serErrors] = useState<Error[]>([])
  const [username, setUsername] = useState<string>("")
  const [password, setPassword] = useState<string>("")

  useEffect(() => {
    const storage = getFromLocalStorage("auth.demo.io")

    if (typeof storage !== "undefined" && storage !== null) {
      setUsername(storage.login.username)
      setPassword(storage.login.password)
    }
  }, [])

  const handleSubmit = async () => {
    addToLocalstorage("auth.demo.io", {
      login: {
        username,
        password,
      },
    })

    const { data, error } = await signinRequest(username, password)
    if (typeof error !== "undefined") {
      serErrors(error)
    }
  }

  return (
    <Form onSubmit={handleSubmit} header="Signin" classNames="mt-16">
      <FormInput
        fieldName="username"
        label="Username:"
        type="text"
        errors={errors}
        onChange={(e) => setUsername(e.target.value)}
        value={username}
      />
      <FormInput
        fieldName="password"
        label="Password:"
        value={password}
        errors={errors}
        onChange={(e) => setPassword(e.target.value)}
        type="password"
      />

      <Button className="mb-5">Log in to your account</Button>
      <Link href="/signup">
        <a className="text-sm text-purpel-500 underline">
          Create a account if you do not have one
        </a>
      </Link>
    </Form>
  )
}

export default SinginForm
