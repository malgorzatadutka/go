package main

import (
	f "03_get_header/functions"
	"fmt"
)

func main() {
	header_type, err := f.GetHeader("https://onet.pl", "Access-Control-Max-Age")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(header_type)
	}
}
