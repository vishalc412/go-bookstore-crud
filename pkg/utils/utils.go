package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	// Read body
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err == nil { // Only proceed if there is NO error
		// Parse JSON only if reading body was successful
		err = json.Unmarshal(body, x)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
