package main

import (
	"WebLog/data"
	"WebLog/manage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/create", manage.Create)
	r.GET("/get", manage.Get)
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", data.Config.Web))
	data.ErrHandle(err)
}

func init() {
	raw, err := os.ReadFile("config.yaml")
	data.ErrHandle(err)
	err = yaml.Unmarshal(raw, &data.Config)
	data.ErrHandle(err)
	data.Data = make(map[string]data.Info)
	go manage.Clean()
	go manage.Listener()
}
