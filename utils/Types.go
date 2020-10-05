// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package utils

import (
	"bytes"
	"net/http"
)

type DumpResponsePayload struct {
	Headers http.Header
	Body    *bytes.Buffer
	Status  int
}
