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
	r.GET("/todos", services.GetTodos)
	// route pour la creation de todo
	r.POST("/todos", services.CreateTodo)

	// lancement du serveur
	r.Run(":8080")
}
