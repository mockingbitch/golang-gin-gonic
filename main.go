package main

import (
    "github.com/gin-gonic/gin"
    "./controllers"
    )

// func main() {
// router := gin.Default()
// api := router.Group("/api")
// {
// api.GET("/ping", func(ctx *gin.Context) {
// ctx.JSON(200, gin.H{
// "message": "pingpong successful",
// })
// })
// }
// router.Run(":9000")
// }

func setupRouter() *gin.Engine {
    r := gin.Default
    r.Static("/public", "./public")

    client := r.Group("/api")
    {
        client.GET("/user/:id", controller.Read)
        client.POST("/user/create", controller.Create)
        client.PATCH("/user/update/:id", controller.Update)
        client.DELETE("/user/delete/:id", controller.Delete)
    }

    return r
}

func main() {
    r := setupRouter()
    r.Run(":9090")
}
