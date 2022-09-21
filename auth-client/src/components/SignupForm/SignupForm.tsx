import { useState } from "react"

import { Error } from "../../types/errors"
import FormInput from "../shared/FormInput"
import { signupRequest } from "./signupRequest"
import Form from "../shared/Form"

const SignupForm: React.FC = () => {
  const [fullName, setFullName] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const [email, setEmail] = useState<string>("")
  const [username, setUsername] = useState<string>("")
  const [errors, setErrors] = useState<Error[]>([])

  const handleSubmit = async () => {
    const { error } = await signupRequest({
      fullName,
      password,
      username,
      email,
    })
    console.log(error)
    if (typeof error !== "undefined") setErrors(error.errors)
  }

  return (
    <Form onSubmit={handleSubmit}>
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
    </Form>
  )
}

export default SignupForm
