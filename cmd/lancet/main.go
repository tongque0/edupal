package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/random"
)

func main() {
	uuid, err := random.UUIdV4()
	if err != nil {
		return
	}
	fmt.Println(uuid)
}
