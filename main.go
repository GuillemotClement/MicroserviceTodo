package main

import (
	"GoToDo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// creation du serveur Gin
	r := gin.Default()

	// definition du dossier static
	r.Static("/static/", "./frontend")

	// ajout de la route pour rendre le frontend
	r.GET("/", services.HandleStatic)

	// definition des routes pour chaque action
	// route affichage des todo
	r.GET("/todos", services.GetTodos)
	// route pour la creation de todo
	r.POST("/todos", services.CreateTodo)
	// route pour update Done
	r.PUT("/todos/:ID", services.ToggleTodo)
	// route pour la supresion
	r.DELETE("/todos/:ID", services.DeleteTodo)

	// lancement du serveur
	r.Run(":8080")
}
