import { Box } from "@mui/system"
import { Input } from "../Input"
import { useSignupFormState } from "./useSignupFormState"

const SignupForm: React.FC = () => {
  const {
    errs,
    setErrs,
    email,
    setEmail,
    userName,
    setUserName,
    password,
    setPassword,
  } = useSignupFormState()

  return (
    <Box
      sx={{
        marginTop: 10,
        padding: 5,
        maxWidth: 400,
        height: "auto",
        backgroundColor: "rgb(32, 32, 32)",
        borderRadius: 10,
      }}
    >
      <Input
        type="email"
        fieldName="email"
        value={email}
        errors={errs}
        onChange={(e) => setEmail(e.target.value)}
      />
      <Input
        type="text"
        fieldName="text"
        value={email}
        errors={errs}
        onChange={(e) => setEmail(e.target.value)}
      />
      <Input
        type="text"
        fieldName="userName"
        value={userName}
        errors={errs}
        onChange={(e) => setUserName(e.target.value)}
      />
    </Box>
  )
}

export default SignupForm
