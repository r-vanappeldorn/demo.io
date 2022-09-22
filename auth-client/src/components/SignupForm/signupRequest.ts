import { ErrorRes } from "@demo.io/utils/dist/types"
import { nextJsApiRequest } from "@demo.io/utils/dist/utils"

export const signupRequest = async (
  fullName: string,
  password: string,
  username: string,
  email: string
) => {
  const { error, data } = await nextJsApiRequest<unknown, ErrorRes>(
    "post",
    { fullName, password, username, email },
    "/api/auth/signup",
    process.env.NEXT_PUBLIC_MAIN_URL,
    process.env.NEXT_PUBLIC_ACCOUNT_SRV
  )

  return {
    error: (error.response?.data as any).errors as Error[],
    data,
  }
}
