import { Container, Box } from "@mui/system"
import type { NextPage } from "next"
import SignupForm from "../components/SignupForm/SignupForm"

const Signup: NextPage = () => {
  return (
    <Container
      fixed
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <SignupForm/>
    </Container>
  )
}

export default Signup
