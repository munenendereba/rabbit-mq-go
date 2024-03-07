package helper

import "log"

/*
	takes an error and error message as input and logs to console
*/
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
