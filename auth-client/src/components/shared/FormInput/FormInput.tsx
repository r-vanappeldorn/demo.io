import React, { FC, useEffect, useState } from "react"
import { Error } from "../../../types/errors"

interface Props {
  fieldName: string
  label: string
  type: "password" | "email" | "text" | "number"
  errors: Error[]
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  value: string
}

const FormInput: FC<Props> = ({
  fieldName,
  label,
  errors,
  type,
  value,
  onChange,
}) => {
  const [error, setError] = useState<Error | undefined>({})

  useEffect(() => {
    if (typeof errors !== "undefined") {
      const filter = errors.filter((err) => err.field === fieldName)
      if (typeof filter[0] !== "undefined") setError(filter[0])
    }
  }, [errors])

  useEffect(() => {
    setError({})
  }, [value])

  return (
    <div className="form-floating mb-6 w-[100%]">
      <input
        type={type}
        className={`form-control block w-full text-base font-normal dark:bg-dark-700 focus:border focus:border-purpel-500 bg-clip-padding rounded transition ease-in-out dark:focus:ring-0 border-transparent dark:focus:border dark:focus:border-dark-100 dark:text-dark-100  ${
          error?.field === fieldName
            ? "border-red-600 border"
            : "dark:border-2 border dark:border-dark-600 border-purpel-500"
        }`}
        id={fieldName}
        placeholder="Password"
        onChange={onChange}
      />
      <label htmlFor={fieldName} className="text-gray-400">
        {label}
      </label>
      {error?.field === fieldName && (
        <p className="text-red-900">{error?.message}</p>
      )}
    </div>
  )
}

export default FormInput
