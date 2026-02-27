package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joseph0x45/surge/internal/db"
	"github.com/joseph0x45/surge/internal/models"
	"github.com/joseph0x45/goutils"
	"github.com/matoous/go-nanoid/v2"
)

func createUser(args []string) int {
	flagSet := flag.NewFlagSet("create-user", flag.ContinueOnError)
	username := flagSet.String("username", "", "The new user's username")
	password := flagSet.String("password", "", "The new user's password")
	flagSet.Parse(args)
	if *username == "" || *password == "" {
		fmt.Fprintln(os.Stderr, "username and password are required")
		return 1
	}
	if len(*password) >= 72 {
		fmt.Fprintln(os.Stderr, "Password is too long.")
		return 1
	}
	conn := db.GetConn(false)
	defer conn.Close()
	exists, err := conn.UsernameExists(*username)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if exists {
		fmt.Println("Username is already taken")
		return 0
	}
	hash, err := goutils.HashPassword(*password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	id, err := gonanoid.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	user := &models.User{
		ID:       id,
		Username: *username,
		Password: hash,
	}
	if err := conn.InsertUser(user); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	log.Println("User created")
	return 0
}
