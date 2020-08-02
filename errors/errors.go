package errors

import "fmt"

//GetError access error
type GetError struct {
}

func (e *GetError) Error() string {
	return fmt.Sprintf("Unable to retrieve object with given conditions")
}

//AddError add error
type AddError struct {
}

func (e *AddError) Error() string {
	return fmt.Sprintf("Unable to add object with given conditions")
}
