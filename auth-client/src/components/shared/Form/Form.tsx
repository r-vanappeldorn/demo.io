import Button from "../Button"
import Logo from "../Logo"

interface Props {
  children: React.ReactNode[] | React.ReactNode[]
  onSubmit(e: React.FormEvent<HTMLFormElement>): void
}

const Form: React.FC<Props> = ({ onSubmit, children }) => (
  <form
    onSubmit={onSubmit}
    className="mx-auto mt-10 dark:bg-dark-800 dark:border-none  bg-white border h-auto pt-10 pb-10 px-10 max-w-[450px] flex flex-col justify-center items-center rounded-xl"
  >
    <Logo />
    <h3 className="text-3xl text-gray-800 mb-5 dark:text-white">Signup</h3>
    {children}

    <Button>Submit</Button>
  </form>
)

export default Form
