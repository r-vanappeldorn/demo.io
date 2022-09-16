import { TextField } from "@mui/material"
import React from "react"
import { Error } from "../../types/errors"
import { useInputState } from "./useInputState"

interface Props {
  type: "email" | "text" | "number"
  fieldName: string
  errors?: Error[]
  onChange: (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => void
  value: string
}

const Input: React.FC<Props> = ({ type, errors, fieldName, onChange, value }) => {
  const { errorMessage } = useInputState({
    type,
    fieldName,
    errors: errors || undefined,
  })

  return (
    <div>
      <TextField
        onChange={onChange}
        id="outlined-basic"
        label="Outlined"
        variant="outlined"
      />
      {errorMessage && <p>{errorMessage}</p>}
    </div>
  )
}

export default Input
