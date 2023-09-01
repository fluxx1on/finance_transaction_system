package logger

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

func GCodeSuite(method string, code codes.Code) string {
	return fmt.Sprintf("%s %s", method, code.String())
}
