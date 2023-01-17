package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var errInvalidRuntimeFormat = errors.New("invalid runtime format")

// Because we want to customize the runtime to be
// something like '32 mins' we use our own marshal
func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedString := strconv.Quote(jsonValue)

	return []byte(quotedString), nil
}

func (r *Runtime) UmarshalJSON(jsonValue []byte) error {

	unquotedJSON, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return errInvalidRuntimeFormat
	}
	// Split the string, isolate the part containing the number
	parts := strings.Split(unquotedJSON, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return errInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return errInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
