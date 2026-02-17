import '@styles/Button.css'

function Button({ children, variant = 'primary', onClick, disabled, ...props }) {
  return (
    <button
      className={`btn btn-${variant}`}
      onClick={onClick}
      disabled={disabled}
      {...props}
    >
      {children}
    </button>
  )
}

export default Button
