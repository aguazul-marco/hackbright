package main

import (
	"github.com/aguazul-marco/hackbright/capstone/api"
	_ "github.com/lib/pq"
)

const (
	port = ":8000"
)

func main() {

	r := api.SetUpRouter()

	r.Run(port)

}
