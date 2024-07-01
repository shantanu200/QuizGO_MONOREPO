package utils

import "fmt"

func RequestLogger(doc interface{}) {
	fmt.Println("Request Params: ", doc)
}
