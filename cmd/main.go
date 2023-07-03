package main

import (
	"fmt"

	"github.com/bdemetris/google-license-cleanup/google"
)

// query users
// https://developers.google.com/admin-sdk/directory/v1/guides/search-users

// user activity api
// https://developers.google.com/admin-sdk/reports/reference/rest/v1/userUsageReport

func main() {
	srv, err := google.NewService()
	if err != nil {
		fmt.Println(err)
	}

	users, err := srv.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(users))
}
