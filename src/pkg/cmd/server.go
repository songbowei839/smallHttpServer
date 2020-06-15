package cmd

import (
	"pkg/server"
)

// Create the number server and run it
func DoNumberServerCmd() {
	n, err := server.CreateServer()
	if err != nil {
		print("Create server failed, terminating")
	}
	n.Run()
}
