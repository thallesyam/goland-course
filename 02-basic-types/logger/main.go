package logger

import (
	"fmt"
	"os"
)

func LoggingError(val *float64) (int, error) {
	return fmt.Fscanln(os.Stdin, val)
}
