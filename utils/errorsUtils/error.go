package errorsUtils

import "log"

func HandleErrorSoft(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func HandleErrorByPanic(err error) {
	if err != nil {
		panic(err)
	}
}
