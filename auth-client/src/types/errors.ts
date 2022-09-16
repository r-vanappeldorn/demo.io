export interface Error {
  message?: string
  field?: string
}

export interface ErrorRes {
  errors: Error[]
}
