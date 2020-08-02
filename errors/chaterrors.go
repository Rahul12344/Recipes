package errors

import "fmt"

//ChatWriteError write to chat error
type ChatWriteError struct {
}

func (e *ChatWriteError) Error() string {
	return fmt.Sprintf("Unable to write to chat")
}

//ChatReadError read from chat error
type ChatReadError struct {
}

func (e *ChatReadError) Error() string {
	return fmt.Sprintf("Unable to read from chat")
}

//UpgradeError read from chat error
type UpgradeError struct {
	Issue error
}

func (e *UpgradeError) Error() string {
	return fmt.Sprintf("Upgrade error; error due to: %s", e.Issue.Error())
}
