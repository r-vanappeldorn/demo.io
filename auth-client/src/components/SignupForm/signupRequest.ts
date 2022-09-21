import axios, { AxiosError } from "axios"
import { ErrorRes } from "../../types/errors"

interface Args {
  fullName: string
  password: string
  username: string
  email: string
}

export const signupRequest = async ({
  fullName,
  password,
  username,
  email,
}: Args) => {
  if (typeof window !== "undefined") {
    try {
      const { data } = await axios.post("http://demo.io/api/auth/signup", {
        fullName,
        email,
        username,
        password,
      })
      return { data }
    } catch (err) {
      return { error: (err as AxiosError).response?.data as ErrorRes }
    }
  }

  try {
    const { data } = await axios.post(
      "http://acount-service-srv/api/auth/signup",
      {
        fullName,
        email,
        username,
        password,
      }
    )
    return { data }
  } catch (err) {
    return { error: (err as AxiosError).response?.data as ErrorRes }
  }
}
