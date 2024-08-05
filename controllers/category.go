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

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
		req    = c.Request
	)

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result = gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
	} else {
		var auth = auth.Auth(uname, pwd)

		// fmt.Println("auth status : ", auth)

		if auth == true {

			categories, err := repository.GetAllCategory(database.DbConnection)

			if err != nil {
				result = gin.H{
					"result": err.Error(),
				}
			} else {
				result = gin.H{
					"result": categories,
				}
			}
		} else {
			result = gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCategory(c *gin.Context) {
	var (
		result   gin.H
		category structs.Category
		req      = c.Request
	)
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result = gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
	} else {

		var auth = auth.Auth(uname, pwd)

		// fmt.Println("auth status : ", auth)

		if auth == true {

			categories, err := repository.GetCategory(database.DbConnection, category)

			if err != nil {
				result = gin.H{
					"result": err.Error(),
				}
			} else {
				result = gin.H{
					"result": categories,
				}
			}
		} else {
			result = gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var (
		category structs.Category
		req      = c.Request
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

		// fmt.Println("auth status : ", auth)

		if auth == true {

			err := c.BindJSON(&category)
			if err != nil {
				panic(err)
			}

			err = repository.InsertCategory(database.DbConnection, category)
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, category)
		} else {
			result := gin.H{
				"error":   "401",
				"message": "Username atau Password salah!",
			}
			c.JSON(http.StatusOK, result)
		}
	}
}

func UpdateCategory(c *gin.Context) {
	var (
		category structs.Category
		req      = c.Request
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result := gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
		c.JSON(http.StatusOK, result)
	} else {

		var auth = auth.Auth(uname, pwd)

		// fmt.Println("auth status : ", auth)

		if auth == true {

			err = repository.UpdateCategory(database.DbConnection, category)
			if err != nil {
				result := gin.H{
					"error":   "400",
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
			} else {
				c.JSON(http.StatusOK, category)
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

func DeleteCategory(c *gin.Context) {
	var (
		category structs.Category
		req      = c.Request
	)
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	uname, pwd, ok := req.BasicAuth()

	if !ok {
		result := gin.H{
			"error":   "401",
			"message": "Auth Kosong!",
		}
		c.JSON(http.StatusOK, result)
	} else {

		var auth = auth.Auth(uname, pwd)

		// fmt.Println("auth status : ", auth)

		if auth == true {

			err := repository.DeleteCategory(database.DbConnection, category)
			if err != nil {
				result := gin.H{
					"error":   "400",
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
			} else {
				c.JSON(http.StatusOK, category)
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
