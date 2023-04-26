package util

import (
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

// Sends a POST request to the given url with the given data
func PostRaw(url string, body map[string]interface{}) error {

	req, _ := sonic.Marshal(body)

	reader := strings.NewReader(string(req))

	_, err := http.Post(url, "application/json", reader)
	if err != nil {
		return err
	}

	return nil
}
