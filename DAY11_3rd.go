package main

import (
	"errors"
	"fmt"
)

var users = map[string]string{
	"lucky": "123",
	"amit":  "passowd122",
}

func login(username, password string) (string, error) {
	if pass, found := users[username]; found {
		if pass == password {
			return "Welcom" + username, nil
		}
		return "", errors.New("Worng password")
	}
	return "", errors.New("User not found")

}
func main() {

	nam, err := login("lucky", "123")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(nam)
	}

	new, err := login("Lipu", "7768")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(new)
	}
}
