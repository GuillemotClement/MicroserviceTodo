package main

import (
	"GoToDo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// creation du serveur Gin
	r := gin.Default()

	// definition des routes pour chaque action
	// route affichage des todo
	r.GET("/todo", services.GetTodos)

	// lancement du serveur
	r.Run(":8080")
}
