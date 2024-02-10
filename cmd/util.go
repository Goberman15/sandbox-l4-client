package cmd

import (
	// "bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func httpRequestor(method string, path string, body io.Reader) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", baseUrl, path), body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	if method == "POST" || method == "PUT" || method == "PATCH" {
		if body == nil {
			return fmt.Errorf("body is required")
		}
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println(string(resBody))

	return nil
}

func printError(err error, model string, action string) {
	fmt.Fprintf(os.Stderr, "Error %s %s: %v", action, model, err)
}
