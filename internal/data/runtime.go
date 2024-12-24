package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// implement a MarshalJSON() method in Runtime type so that it satisfies
// json.Marshaler interface. This should return the JSON-encoded value for the
// movie runtime (it will return in the format "<runtime> mins")
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
