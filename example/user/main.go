package main

import (
	"errors"
	"fmt"
	userSDK "github.com/TitoCarpio/go_course_sdk/user"
	"os"
)

func main() {
	userTrans := userSDK.NewHttpClient("http://localhost:8001", "")
	user, err := userTrans.Get("4f843029-d118-4e5b-9292-1e542648e8d7")

	if err != nil {
		if errors.As(err, &userSDK.ErrNotFound{}) {
			fmt.Println("User not found", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("User: ", user)
}
