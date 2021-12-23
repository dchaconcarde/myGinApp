package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var users []User
var token string = "normalicenLaMayonesaConFideosYTuco"

func main() {
	r := gin.Default()
	r.GET("/users", GetAll)
	r.POST("/users", saveUser)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Sigo con esto")
}

type User struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre" binding:"required"`
	Apellido    string  `json:"apellido" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	Edad        int     `json:"edad" binding:"required"`
	Altura      float64 `json:"altura" binding:"required"`
	Activo      bool    `json:"activo" binding:"required"`
	FechaCreado string  `json:"fechaCreado"`
}

func saveUser(c *gin.Context) {
	var user User

	tokenFromHead := c.GetHeader("token")
	if token != tokenFromHead {
		c.JSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = len(users)
	user.FechaCreado = time.Now().String()
	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{"Usuarios actuales": users})

}

func GetAll(c *gin.Context) {
	tokenFromHead := c.GetHeader("token")
	if token != tokenFromHead {
		c.JSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Usuarios actuales": users})

}
