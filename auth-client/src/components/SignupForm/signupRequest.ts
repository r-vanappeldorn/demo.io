import { ErrorRes } from "@demo.io/utils/dist/types"
import { nextJsApiRequest } from "@demo.io/utils/dist/utils"

export const signupRequest = async (
  fullName: string,
  password: string,
  username: string,
  email: string
) => {
  const { error } = await nextJsApiRequest<unknown, ErrorRes>(
    "post",
    { fullName, password, username, email },
    "/api/auth/signup"
  )

  return {
    error: (error.response?.data as any).errors as ErrorRes,
  }
}
