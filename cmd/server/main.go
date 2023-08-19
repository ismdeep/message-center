package main

import (
	"github.com/ismdeep/message-center/app/server/rest"
)

func main() {
	if err := rest.Run(); err != nil {
		panic(err)
	}
}
