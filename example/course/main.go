package main

import (
	"errors"
	"fmt"
	courseSDK "github.com/TitoCarpio/go_course_sdk/course"
	"os"
)

func main() {
	userTrans := courseSDK.NewHttpClient("http://localhost:8002", "")
	user, err := userTrans.Get("1c5d55c3-d09f-4ddf-a032-b7dcfdf71b11")

	if err != nil {
		if errors.As(err, &courseSDK.ErrNotFound{}) {
			fmt.Println("User not found", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Course: ", user)
}
