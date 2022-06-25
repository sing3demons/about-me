package main

import (
	"fmt"
)

type aboutMe struct {
	pronouns  []string
	languages []string
	hobbies   []string
}

type contact struct {
	Facebook  string
	Messenger string
	Line      string
	Telephone string
	Email     []string
	Github    string
}

func main() {
	fmt.Println("Hello, I'm Kumneung!👋")

	sing3demons := aboutMe{
		pronouns:  []string{"He", "Him"},
		languages: []string{"Go", "JavaScript", "C#", "Java"},
		hobbies:   []string{"Watching", "Sleeping"},
	}

	fmt.Printf("Pronouns: %v\n", sing3demons.pronouns)
	fmt.Printf("Languages I Know: %v\n", sing3demons.languages)
	fmt.Printf("My Hobbies: %v\n", sing3demons.hobbies)

	c := contact{
		Facebook:  "facebook.com/sing3demons",
		Messenger: "https://m.me/sing3demons",
		Line:      "sing3demons",
		Telephone: "+66974799593",
		Email:     []string{"sing3demons@live.com", "kp.sing3demons@gmail.com"},
		Github:    "github.com/sing3demons",
	}
	s := map[string]any{"Line": c.Line, "Email": c.Email, "Messenger": c.Messenger, "Facebook": c.Facebook, "Telephone": c.Telephone, "Github": c.Github}
	fmt.Printf("Contact: %v\n", s)
}
