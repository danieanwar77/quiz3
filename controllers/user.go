package controllers

import (
	"net/http"
	"quiz3/database"
	repository "quiz3/repositories"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	var (
		result gin.H
		user   structs.Users
	)
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id

	users, err := repository.GetUser(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.Users

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user structs.Users
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = id

	err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user structs.Users
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id
	err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

// func CheckAuth() {
// 	// konfigurasi server
// 	server := &http.Server{
// 		Addr: ":8000",
// 	}

// 	// routing
// 	http.Handle("/", auth.Auth(http.HandlerFunc(auth.CheckAuth)))

// 	// jalankan server
// 	fmt.Println("server running at http://localhost:800")
// 	err := server.ListenAndServe()
// 	fmt.Println("error server", err)
// }
