package functions

import (
	"fmt"
	"net/http"
)

func GetHeader(url string, header string) (string, error) {
	response, err := http.Get(url)
	// fmt.Print(response)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	header_type := response.Header.Get(header)
	if header_type == "" {
		return "", fmt.Errorf("Header not found!")
	}
	return header_type, nil
}
