import { Error } from "../../types/errors"

export const getMessageForFieldName = (errors: Error[], fieldName: string) => {
  return errors.filter((err) => {
    return err.field === fieldName
  })[0]["message"]
}
