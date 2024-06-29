package errorsUtils

import (
	"fmt"
)

func HandleErrorSoft(err error) {
	if err != nil {
		fmt.Printf("Soft error occured: %v\n", err)
	}
}
func HandleErrorByPanic(err error) {
	if err != nil {
		panic(err)
	}
}
