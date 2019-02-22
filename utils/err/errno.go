package err

import "fmt"

type Errno struct {
	Code    int
	Message string
}

// func Message() string {

// }

func (err *Errno) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s", err.Code, err.Message)
}
