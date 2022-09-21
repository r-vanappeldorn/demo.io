import React, { FC } from "react"

interface Props {
  children?: React.ReactNode | React.ReactNode[]
  onClick?: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
  className?: string
}

const Button: FC<Props> = ({ children, onClick, className }) => {
  return (
    <button
      className={`w-[100%] bg-purpel-500  rounded py-2.5 hover:bg-purpel-600  text-gray-200 ${
        className && className
      }`}
      onClick={onClick && onClick}
    >
      {children}
    </button>
  )
}

export default Button
