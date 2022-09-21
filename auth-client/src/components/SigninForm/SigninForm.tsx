import { Button, Form, FormInput } from "@demo.io/utils/dist/components"
import { Error } from "@demo.io/utils/dist/types"
import Link from "next/link"
import { FC, useState } from "react"

const SinginForm: FC = () => {
  const [errors, serErrors] = useState<Error[]>([])
  const [username, setUsername] = useState<string>("")
  const [password, setPassword] = useState<string>("")

  const handleSubmit = async () => {
    console.log(username, password)
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
        <a className="text-sm  text-purpel-500 underline" href="/signup">
          Create a account if you don't have one
        </a>
      </Link>
    </Form>
  )
}

export default SinginForm
