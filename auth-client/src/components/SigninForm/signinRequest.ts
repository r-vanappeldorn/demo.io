import { Error, ErrorRes } from "@demo.io/utils/dist/types"
import { nextJsApiRequest } from "@demo.io/utils/dist/utils"

export const signinRequest = async (username: string, password: string) => {
  const { error, data } = await nextJsApiRequest<unknown, ErrorRes>(
    "post",
    { password, username },
    "/api/auth/signin",
    process.env.NEXT_PUBLIC_MAIN_URL,
    process.env.NEXT_PUBLIC_ACCOUNT_SRV
  )


  return {
    error: (error.response?.data as any).errors as Error[],
    data,
  }
}
