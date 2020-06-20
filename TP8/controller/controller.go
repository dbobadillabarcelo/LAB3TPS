package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Sumar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if (a == "" || b == "") {
		c.String(400, "Missing params")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if (error1 != nil || error2 != nil) {
		c.String(400, "Invalid params")
		return
	}
	result := num1 + num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}

func Restar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if (a == "" || b == "") {
		c.String(400, "Missing params")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if (error1 != nil || error2 != nil) {
		c.String(400, "Invalid params")
		return
	}
	result := num1 - num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}


func Multiplicar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if (a == "" || b == "") {
		c.String(400, "Missing params")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if (error1 != nil || error2 != nil) {
		c.String(400, "Invalid params")
		return
	}
	result := num1 * num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}


func Dividir(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if (a == "" || b == "") {
		c.String(400, "Missing params")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if (error1 != nil || error2 != nil) {
		c.String(400, "Invalid params")
		return
	}
	if (num2 == 0) {
		c.String(400, "Invalid operation")
		return
	}
	result := num1 / num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}