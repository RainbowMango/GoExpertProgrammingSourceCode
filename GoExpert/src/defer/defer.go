package deferexample

import "fmt"

func NoPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
	}
}

func Dived(n int) {
	defer NoPanic()

	fmt.Println(1 / n)
}

func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}

func UpdateTable() {
	defer func() {
		if IsPanic() {
			// Rollback transaction
		} else {
			// Commit transaction
		}
	}()

	// Database CRUD operation...
}
