package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Because we want to customize the runtime to be
// something like '32 mins' we use our own marshal
func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedString := strconv.Quote(jsonValue)

	return []byte(quotedString), nil
}
