package models

// creation de la struct d'une todo
type Todo struct {
	ID    string // Id unique permettant de pointer vers cette todo
	Title string //Visible pour l'user
	Done  bool   //Permet de gerer le statut de la tache (finis ou non)
}
