package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Choose a mode: \n 1. Regular \n 2. Hard (Troll Alert)")
	var c string
	fmt.Scanln(&c)
	c = strings.ToLower(c)
	if c == "regular" {

		opt := []string{"rock", "paper", "scissor"}
		choose := opt[rand.Intn(len(opt))]
		fmt.Println("Enter Rock, Paper or Scissor")
		var uc string
		fmt.Scanln(&uc)
		uc = strings.ToLower(uc)

		if choose == "rock" && uc == "scissor" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nComputer wins")
		}
		if choose == "rock" && uc == "rock" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nIt is a tie")
		}
		if choose == "rock" && uc == "paper" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nUser wins")

		}
		if choose == "paper" && uc == "scissor" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nUser wins")
		}
		if choose == "paper" && uc == "paper" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nIt is a tie")
		}

		if choose == "paper" && uc == "rock" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nComputer wins")

		}

		if choose == "paper" && uc == "scissor" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nUser wins")
		}
		if choose == "scissor" && uc == "scissor" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nIt is a tie")
		}
		if choose == "scissor" && uc == "rock" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nUser wins")
		}

		if choose == "scissor" && uc == "paper" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Printf("\nComputer's choice: %v", choose)
			fmt.Println("\nComputer wins")
		}

	}
	if c == "hard" {
		fmt.Println("Enter Rock, Paper or Scissor")
		var uc string
		fmt.Scanln(&uc)
		uc = strings.ToLower(uc)
		if uc == "rock" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Println("\nComputer's choice: paper")
			fmt.Println("\nComputer wins ")
		}
		if uc == "paper" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Println("\nComputer's choice: scissor")
			fmt.Println("\nComputer wins")
		}
		if uc == "scissor" {
			fmt.Printf("User's choice: %v", uc)
			fmt.Println("\nComputer's choice: rock")
			fmt.Println("\nComputer wins")
		}

	}
}
