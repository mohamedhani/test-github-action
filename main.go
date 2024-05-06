// main.go
package main

import (
	"github.com/mohamedhani/test-github-action/router"
)

func main() {
	// fmt.Println("welcome")
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		panic(err)

	}

}
