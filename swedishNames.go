package swedishNames

import (
	"io"
	"net/http"
)

// GetFnamn fetches data from a specific SCB API endpoint and returns the response body as a string.
func GetFnamn() string {
	resp, _ := http.Get("https://api.scb.se/OV0104/v1/doris/sv/ssd/START/BE/BE0001/BE0001G/BE0001T06AR")
	defer resp.Body.Close()
	x, _ := io.ReadAll(resp.Body)

	return string(x)
}
