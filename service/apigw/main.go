package main

import (
	"github.com/feitianlove/FIleStore/service/apigw/route"
)

func main() {
	r := route.Router()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
