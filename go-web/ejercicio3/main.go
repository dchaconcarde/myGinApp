package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/users", GetAll)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email string `json:"email"`
	Edad int `json:"edad"`
	Altura float64 `json:"altura"`
	Activo bool `json:"activo"`
	FechaCreado string `json:"fechaCreado"`
}

func readJson() Users{

	jsonFile, err := os.Open("/Users/dchaconcarde/go/src/github.com/dchaconcarde/myGinApp/go-web/users.json")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal([]byte(byteValue), &users)
	
	return users
}

func GetAll(c *gin.Context){
	users := readJson()
	i:=0
	for _, user := range users.Users {
		c.String(http.StatusOK, fmt.Sprintf("%#v \n", user))
		i++
	}
}