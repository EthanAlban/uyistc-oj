package models

import (
	"fmt"
	"testing"
)

func TestFIndUsersByUsername(t *testing.T) {
	user, err := FIndUsersByUsername("et")
	fmt.Println(user, err)
}

func TestAllUsers(t *testing.T) {
	users, err := AllUsers()
	if err == nil {
		for i := 0; i < len(users); i++ {
			fmt.Println(users[i])
		}
	}
}
