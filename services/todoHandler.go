package services

import (
	"GoToDo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
func CreateTodo(c *gin.Context) {
	// on veux creer une nouvelle todo
	// on lui ajoute en type notre struct Todo
	var newTodo models.Todo

	// on verifie que la conversion est reussie
	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		// si j'ai une erreur, je viens generer un code erreur
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": err})
		// je return pour stoper la requete
		return
	}

	// si la conversion a reussis alors je continue la creation d'un todo
	// je creer un nouvel ID pour ma todo
	newTodo.ID = uuid.New().String()
	// j'ajoute ensuite dans le todo
	todos = append(todos, newTodo)
	// on retourne la todo creer avec un code 201
	c.JSON(http.StatusCreated, newTodo)
}

// 4 - Modifier une todo => PUT
// on souhaite venir modifier le status de la task
func ToggleTodo(c *gin.Context) {
	// on recupere l'id depuis le parametre passer via url
	id := c.Param("ID")
	// on viens iterer sur les todo pour trouver celle qui correspond a l'id
	// todo correspond a la todo iterer
	for i, todo := range todos {
		// si id de la todo iterer correspond a id recup depuis la requete
		if todo.ID == id {
			// on viens modifier la todo present dans la liste
			// on viens echanger la valeur de Done pour indiquer que la tache est finis ou non
			todos[i].Done = !todos[i].Done
			// on retourne la tache modifier et le code 200
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	// si la todo n'est pas trouver on retourne un message d'erreur
	c.JSON(http.StatusNotFound, gin.H{"Erreur": "Todo not find"})
}

// 5 - Supprimer une todo => DELETE
