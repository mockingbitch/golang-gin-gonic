package controllers

import (
    "github.com/gin-gonic/gin"
    "../database"
)

type User struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Content string `json:"body"`
}

func Read (c * gin.Context) {
    db := database.DBConn()
    rows, err := db.Query("SELECT id, name, age FROM users WHERE id = " + c.Param("id"))
    if err != nil {
        c.JSON(500, gin.H {
            "messages": "User not found",
        })
    }

    post := User{}

    for rows.Next() {
        var id int
        var name, age string

        err = rows.Scan(&id, &name, &age)
        if err != nil {
            panic(err.Error())
        }

        user.Id = id
        user.Name = name
        user.Age = age
    }

    c.JSON(200, user)
    defer db.Close()
}
