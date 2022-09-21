import { ErrorRes } from "@demo.io/utils/dist/types"
import { nextJsApiRequest } from "@demo.io/utils/dist/utils"

export const signinRequest = async (
  fullName: string,
  password: string,
  username: string,
  email: string
) => {
  const { error, data } = await nextJsApiRequest<unknown, ErrorRes>(
    "post",
    { fullName, password, username, email },
    "/api/auth/signin"
  )

  console.log(data)

  return {
    error: (error.response?.data as any).errors as ErrorRes,
  }
}
