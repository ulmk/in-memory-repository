package main

import (
	"fmt"

	"github.com/ulmk/in-memory-repository/repo"
)

func main() {
	repository := &repo.InMemoryRepo{}
	user1 := &repo.User{
		Name:     "Meg",
		Email:    "meg@meg.com",
		Password: "password",
	}

	user2 := &repo.User{
		Name:     "Ken",
		Email:    "ken@ken.com",
		Password: "password",
	}

	repository.Create(user1)
	fmt.Println(repository.GetAll())

	repository.Create(user2)
	fmt.Println(repository.GetAll())

	user2.Name = "Jack"
	repository.Update(user2)
	fmt.Println(repository.GetAll())

	repository.Delete(user2.ID)
	fmt.Println(repository.GetAll())

	// users, _ := repository.GetAll()
	// for _, user := range users {
	// 	fmt.Println(user)
	// }
}
