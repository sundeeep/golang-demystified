package main

import (
	"fmt"
	"golang-demystified/models"
)

func main() {
	fmt.Println("hello, world!")

	secondaryEmails := []string{"sandeepkuma0100110@gmail.com", "indiandevgarrage@gmail.com"}

	newUser := models.CreateNewUser(369, "Sundeeep", "Dasari", "sundeeep", "Sundeep369", "6305309369", "sundeeepdev@gmail.com", secondaryEmails, "SUPER_ADMIN")

	newUser.Role = "ADMIN"

	fmt.Println(newUser.Role)
}
