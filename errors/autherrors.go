package errors

import "fmt"

//LoginError login error wrapper
type LoginError struct {
	Issue error
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error; error due to: %s", e.Issue.Error())
}

//SignupError singup error wrapper
type SignupError struct {
	Issue error
}

func (e *SignupError) Error() string {
	return fmt.Sprintf("Signup error; error due to: %s", e.Issue.Error())
}

//InvalidTokenError invalid token
type InvalidTokenError struct {
	token string
}

func (e *InvalidTokenError) Error() string {
	return fmt.Sprintf("Invalid token: %s", e.token)
}
