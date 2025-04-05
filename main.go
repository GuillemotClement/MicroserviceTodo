package main

import "github.com/gin-gonic/gin"

func main() {
	// creation du serveur Gin
	r := gin.Default()

	// definition des routes pour chaque action
	// 1 - Generation d'un statique => GET
	// 2 - Montrer les todo => GET
	// 3 - Creation d'une todo => POST
	// 4 - Modifier une todo => PUT
	// 5 - Supprimer une todo => DELETE

	// lancement du serveur
	r.Run(":8080")
}
