package main

import (
	"fmt"
	"github.com/RakhimovAns/sportgrum/pkg/handlers"
	"github.com/RakhimovAns/sportgrum/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := utils.ReadConfig()
	if err != nil {
		fmt.Printf("Read config file error: %w", err)
	}
	r := gin.Default()
	//r.GET("/ping", handlers.Pong)
	r.GET("/register", handlers.Register)
	fmt.Println(cfg.Port)
	r.Run(cfg.Port)

}
