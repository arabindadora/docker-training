package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

type Payload struct {
    Input string `json:"input"`
}

func sha256Handler(c *gin.Context) {
    var p Payload

    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"result": GetSHA256(p.Input)})
}

func base64Handler(c *gin.Context) {
    var p Payload

    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"result": GetBase64(p.Input)})
}

