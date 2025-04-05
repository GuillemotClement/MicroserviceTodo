const apiURL = "/todos";

//Obtenir
async function fetchTodos() {
  // on fait un call vers l'url
  const response = await fetch(apiURL);
  // on recupere les data du json
  const todos = await response.json();
  // on creer un nouvel element
  const todoList = document.querySelector("#listTodo");

  // on nettoie la todoList
  todoList.innerHTML = "";

  // on boucle sur les todos
  todos.forEach((todo) => {
    // on creer le li
    const li = document.createElement("li");
    // on creer le html de la todo
    li.innerHTML = `
      <span>${todo.Title} - ${todo.Done ? "Terminer" : "En cours"}</span>
      <div>
        <button class="toggleTask">${todo.Done ? "Fait" : "Non fait"}</button>
        <buttom class="deleteBtn">Supprimer</button>
      </div>
      `;

    // on recupere les deux boutons
    const toggleBtn = li.querySelector(".toggleTask");
    const deleteBtn = li.querySelector(".deleteBtn");

    // ajout des evenements
    toggleBtn.addEventListener("click", () => {});
    deleteBtn.addEventListener("click", () => {});

    // ajout des elemts dans le HTML
    todoList.appendChild(li);
  });
}
//Ajouter
async function addTask() {
  // recuperation du input
  const newInput = document.querySelector("#newTask");
  // on recupere les donnes de l'input
  const title = newInput.value.trim();
  // on verifie si title est vide
  if (title == "") {
    alert("Une valeur est attendue");
    return;
  }
  // on fait le call pour envoyer la todo au back
  await fetch(apiURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title, completed: false }),
  });
  // on vide le input
  newInput.value = "";
  // on recupere les nouvelles todos
  fetchTodos();
}
//Modifier

//Delete
