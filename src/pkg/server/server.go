package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "net/http"
	"pkg/api"
)

type numberServer struct {
	endpoint string
	router   *gin.Engine
}

type NumberServer interface {
	Run()
}

func CreateServer() (NumberServer, error) {
	// create gin api router
	r := gin.New()
	zapis.Bind(r)
	return &numberServer{
		//endpoint: viper.GetString("spec.server.address"),
		endpoint: "127.0.0.1:11400",
		router:   r,
	}, nil
}

// start the http server.
func (n *numberServer) Run() {
	fmt.Printf("endpoint %s\n", n.endpoint)
	err := n.router.Run(n.endpoint)
	if err != nil {
		print("Start API gateway failed.")
	}

}
