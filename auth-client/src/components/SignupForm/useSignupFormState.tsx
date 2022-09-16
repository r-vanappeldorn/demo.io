import React, { useEffect, useState } from "react"
import { Error } from "../../types/errors"

const handler = (
  e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  dispatch: React.Dispatch<string>
) => {
  dispatch(e.target.value)
}

export const useSignupFormState = () => {
  const [errs, setErrs] = useState<Error[]>([])
  const [email, setEmail] = useState<string>("")
  const [userName, setUserName] = useState<string>("")
  const [password, setPassword] = useState<string>("")

  useEffect(() => {}, [errs])

  return {
    errs,
    setErrs,
    userName,
    handleUsername: (
      e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => handler(e, setUserName),
    email,
    handleEmail: (
      e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => handler(e, setEmail),
    password,
    setPassword,
  }
}
