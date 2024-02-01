package utils

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// The function "HandleError" logs an error message along with the current time, file name, line
// number, and function name if an error occurs.
func HandleError(message string, err error) {
	if err != nil {
		now := time.Now()
		formatTime := now.Format(time.UnixDate)
		file, line, functionName := trace()
		filename := strings.Split(file, "/")
		errorMessage := fmt.Sprintf("***** %s: %v on line %s in %s at file %s *****", message, err, strconv.Itoa(line), functionName, filename[len(filename)-1])
		errorMessageWithFormatTime := formatTime + ": " + errorMessage + "\n"
		log.Println(errorMessageWithFormatTime)
	}
}

// The trace function returns the file name, line number, and function name of the caller.
func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}
