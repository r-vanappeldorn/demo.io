import { FormEvent, useState } from "react"

import { Button, Form, FormInput } from "@demo.io/utils/dist/components"
import { Error } from "@demo.io/utils/dist/types"
import { signupRequest } from "./signupRequest"
import Link from "next/link"

const SignupForm: React.FC = () => {
  const [fullName, setFullName] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const [email, setEmail] = useState<string>("")
  const [username, setUsername] = useState<string>("")
  const [errors, setErrors] = useState<Error[]>([])

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    const { error } = await signupRequest(fullName, password, username, email)
    console.log(error)
    if (typeof error !== "undefined") setErrors(error.errors)
  }

  return (
    <Form classNames="mt-16" onSubmit={handleSubmit} header="Signup">
      <FormInput
        value={fullName}
        fieldName="fullname"
        type="text"
        errors={errors}
        label="Full name:"
        onChange={(e) => setFullName(e.target.value)}
      />
      <FormInput
        value={email}
        fieldName="email"
        type="email"
        errors={errors}
        label="Email:"
        onChange={(e) => setEmail(e.target.value)}
      />
      <FormInput
        value={username}
        fieldName="username"
        type="text"
        errors={errors}
        label="Username:"
        onChange={(e) => setUsername(e.target.value)}
      />
      <FormInput
        value={password}
        fieldName="password"
        type="password"
        errors={errors}
        label="Password:"
        onChange={(e) => setPassword(e.target.value)}
      />
      <Button className="mb-5">Create your account</Button>
      <Link href="/signin">
        <a className="text-sm text-purpel-500 underline" href="/singin">
          Signin if you already have a account
        </a>
      </Link>
    </Form>
  )
}

export default SignupForm
