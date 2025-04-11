package cmd

import "github.com/private-project-pp/product-rpc-service/interfaces"

func StartServer() {
	if err := interfaces.Container(); err != nil {
		panic(err)
	}
}
