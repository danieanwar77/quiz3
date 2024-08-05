package controllers

import (
	"net/http"
	"quiz3/auth"
	"quiz3/database"
	repository "quiz3/repositories"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var (
		book structs.Book
		req  = c.Request
	)

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result := gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
		c.JSON(http.StatusOK, result)
	} else {

		var auth = auth.Auth(uname, pwd)

		if auth == true {

			err := c.BindJSON(&book)
			if err != nil {
				panic(err)
			}

			err = repository.InsertBook(database.DbConnection, book)
			if err != nil {
				result := gin.H{
					"error":   "400",
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
			} else {
				c.JSON(http.StatusOK, book)
			}
		} else {
			result := gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}
			c.JSON(http.StatusOK, result)
		}
	}
}

func UpdateBook(c *gin.Context) {
	var (
		book structs.Book
		req  = c.Request
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = id

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result := gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
		c.JSON(http.StatusOK, result)
	} else {

		var auth = auth.Auth(uname, pwd)

		if auth == true {

			err = repository.UpdateBook(database.DbConnection, book)
			if err != nil {
				result := gin.H{
					"error":   "400",
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
			} else {
				c.JSON(http.StatusOK, book)
			}

		} else {
			result := gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}

			c.JSON(http.StatusOK, result)
		}
	}
}

func DeleteBook(c *gin.Context) {
	var (
		book structs.Book
		req  = c.Request
	)

	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result := gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
		c.JSON(http.StatusOK, result)
	} else {

		var auth = auth.Auth(uname, pwd)

		if auth == true {

			err := repository.DeleteBook(database.DbConnection, book)
			if err != nil {
				result := gin.H{
					"error":   "400",
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
			} else {
				c.JSON(http.StatusOK, book)
			}

		} else {
			result := gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}

			c.JSON(http.StatusBadRequest, result)
		}
	}
}
