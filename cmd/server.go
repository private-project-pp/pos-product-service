package cmd

import "github.com/private-project-pp/user-rpc-service/interfaces"

func StartServer() {
	if err := interfaces.Container(); err != nil {
		panic(err)
	}
}
