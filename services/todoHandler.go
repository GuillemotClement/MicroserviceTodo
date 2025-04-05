package services

import (
	"GoToDo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// stockage des todo
// on utilise un slice de Todo et on l'initialize vide
// on utilise le struct definis dans les models
var todos = []models.Todo{}

// 1 - Generation d'un statique => GET

// 2 - Montrer les todo => GET
func GetTodos(c *gin.Context) {
	// on viens retourner les differentes todo -> proviens de la DB
	// ici on utilise du JSON car on ne stocke pas les donnees.
	// ici on retourne le sclice de todo qui contiens toutes les todo creer.
	// on peut passer le code ou le http.StatusOK pour 200
	c.JSON(http.StatusOK, todos)
}

// 3 - Creation d'une todo => POST

// 4 - Modifier une todo => PUT

// 5 - Supprimer une todo => DELETE
