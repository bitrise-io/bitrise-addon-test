package utils

import "log"

// LogExceptionIfErr ...
func LogExceptionIfErr(err error) {
	if err != nil {
		log.Printf(" [!] Exception: %+v", err)
	}
}

// LogExceptionIfErrWithContext ...
func LogExceptionIfErrWithContext(err error, exceptionContext string) {
	if err != nil {
		log.Printf(" [!] Exception: %s: %+v", exceptionContext, err)
	}
}
