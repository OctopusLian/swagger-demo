/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-16 16:40:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-16 17:01:15
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"swagger-demo/api"

	"github.com/gin-gonic/gin"

	// This line is necessary for go-swagger to find your docs!
	_ "swagger-demo/docs"
)

var users []*api.User

func main() {
	r := gin.Default()
	r.POST("/users", Create)
	r.GET("/users/:name", Get)
	log.Fatal(r.Run(":5555"))
}

// Create create a user in memory.
func Create(c *gin.Context) {
	var user api.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}
	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}
	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

// Get return the detail information for a user.
func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})
}
