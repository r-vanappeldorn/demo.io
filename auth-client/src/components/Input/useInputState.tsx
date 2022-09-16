import { useEffect, useState } from "react"
import { Error } from "../../types/errors"
import { getMessageForFieldName } from "./helpers"

type Params = {
  type: "email" | "text" | "number"
  fieldName: string
  errors?: Error[]
}

export const useInputState = ({ type, fieldName, errors }: Params) => {
  const [errorMessage, setErrorMessage] = useState<string | undefined>(
    undefined
  )

  useEffect(() => {
    if (typeof errors !== "undefined") 
      setErrorMessage(getMessageForFieldName(errors, fieldName))
  }, [errors])

  return { errorMessage }
}
