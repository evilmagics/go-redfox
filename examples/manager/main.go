package main

import (
	"encoding/json"
	"fmt"

	"github.com/evilmagics/go-redfox"
)

func main() {
	manager := redfox.NewManager[string]()

	// Add many exceptions
	// Using Set function carefully, all registered exceptions will be overwritten
	manager.Set(map[string]redfox.Exception[string]{
		"USERNAME_REQUIRED": redfox.New("USERNAME_REQUIRED", "username must filled"),
		"PASSWORD_REQUIRED": redfox.New("PASSWORD_REQUIRED", "password must filled"),
	})

	// Add multiple exceptions
	manager.AddAll(
		redfox.New("AUTHORIZATION_INVALID", "authorization invalid"),
		redfox.New("SIGNATURE_INVALID", "signature invalid"),
	)

	// Add a new error template
	manager.Add(redfox.New("SERVER_ERROR", "internal server error"))

	// Add safe exceptions make sure that new exception will not overwrite existing exceptions
	err := manager.SafeAdd(redfox.New("DATABASE_ERROR", "database not connected"))
	if err != nil {
		panic(err)
	}

	for _, v := range manager.GetAll() {
		j, err := json.Marshal(v.View())
		if err != nil {
			panic(err)
		}

		fmt.Println(string(j))
	}

}
