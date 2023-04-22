package main

import (
	"go_inven_ctrl/delivery"
	// "test/repository"

	_ "github.com/lib/pq"
)

// "test/delivery"

func main() {
	delivery.Exec()
	// repository.Exec2()
}
