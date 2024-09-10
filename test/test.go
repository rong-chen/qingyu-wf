package main

import (
	"fmt"
	"qingyu-wf/utils"
)

var toke = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWMwNzRjNmMtNmNjNy0xMWVmLWJmYTUtMDQ3YzE2NTMzNDQ0IiwiaXNzIjoicWlueXUiLCJleHAiOjE3MjYwMzkxMTEsImlhdCI6MTcyNTk1MjcxMX0.jRlEFsrIwudrRKE2vaa9lCkYtZgSEApljdB-cOZeLMY"

func main() {
	jwt, err := utils.ParseJWT(toke)
	if err != nil {
		return
	}
	fmt.Println(jwt)
}
