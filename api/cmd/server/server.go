package main

import (
	"github.com/smartlog/api/cmd/builder"
)

func main() {
	apiDependencies, cancel, err := builder.Initialize()
	if err != nil {
		panic(err)
	}

	proto_clients.RegisterSmartLogV1Server()
}
