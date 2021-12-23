package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	users := readJson()

	fmt.Println(users)

	r := gin.Default()
	r.GET("/message", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola %s", users.Users[0].Nombre)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Email       string  `json:"email"`
	Edad        int     `json:"edad"`
	Altura      float64 `json:"altura"`
	Activo      bool    `json:"activo"`
	FechaCreado string  `json:"fechaCreado"`
}

func readJson() Users {

	jsonFile, err := os.Open("/Users/dchaconcarde/go/src/github.com/dchaconcarde/myGinApp/go-web-clase1-TM/users.json")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(byteValue, &users)

	return users
}
