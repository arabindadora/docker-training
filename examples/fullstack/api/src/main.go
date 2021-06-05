package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"*"}
    // OPTIONS method for axios
    corsConfig.AddAllowMethods("OPTIONS")
    // register the middleware
    router.Use(cors.New(corsConfig))

    router.POST("/hash/sha256", sha256Handler)
    router.POST("/encode/base64", base64Handler)
    router.Run(":8000")
}
