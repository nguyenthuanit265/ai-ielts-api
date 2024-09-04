package main

import (
	"main/server"

	_ "main/docs"
)

// @title           Document API
// @version         1.0
// @description     This is a document for API.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9092
// @BasePath  /api/v1

func main() {
	cm := server.CommandManager{}

	cm.AddCommand(server.Command{
		Script: "",
		Desc:   "Run backend server",
		ExecuteFunc: func(args []string) {
			apiServer := server.Initialize()
			apiServer.RunCronJob()
			apiServer.Start()
		},
	})

	cm.Execute()
}
