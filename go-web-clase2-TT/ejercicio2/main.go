package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", GetAll)
	r.GET("/users/activeUsers", filteredUsers)
	r.GET("/users/:id", findUserById)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID          int     `form: "id" json:"id"`
	Nombre      string  `form: "nombre" json:"nombre"`
	Apellido    string  `form: "apellido" json:"apellido"`
	Email       string  `form: "email" json:"email"`
	Edad        int     `form: "edad" json:"edad"`
	Altura      float64 `form: "altura" json:"altura"`
	Activo      bool    `form: "activo" json:"activo"`
	FechaCreado string  `form: "fechaCreado" json:"fechaCreado"`
}

func readJson() Users {

	jsonFile, err := ioutil.ReadFile("/Users/dchaconcarde/go/src/github.com/dchaconcarde/myGinApp/go-web-clase2-TT/users.json")

	if err != nil {
		log.Fatal(err)
	}

	var users Users
	json.Unmarshal(jsonFile, &users)

	return users
}

func GetAll(c *gin.Context) {
	users := readJson()
	i := 0
	for _, user := range users.Users {
		c.JSON(http.StatusOK, user)
		i++
	}
}

func filteredUsers(c *gin.Context) {
	var user []User
	var filteredUsers []User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, us := range user {
		if !us.Activo {
			c.JSON(http.StatusUnauthorized, gin.H{us.Nombre + " " + us.Apellido: "inactivo"})
			continue
		}
		filteredUsers = append(filteredUsers, us)
	}

	c.JSON(http.StatusOK, gin.H{"Usuarios Activos": filteredUsers})

}

func findUserById(c *gin.Context) {
	var user []User
	userId := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, us := range user {
		intUserId := 0
		_, err := fmt.Sscan(userId, &intUserId)
		if err != nil {
			log.Fatal(err)
		}
		if us.ID == intUserId {
			c.JSON(http.StatusOK, gin.H{"Usuario encontrado": us})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "usuario no encontrado"})

}
