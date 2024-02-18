const ui = {
  inputElement: document.getElementById("title"),
  addButtonElement: document.getElementById("add-btn"),
  tableBody: document.querySelector("tbody"),
};

// function to fetch todos
async function fetchTodos() {
  const todos = await fetch("http://localhost:5000/todos");

  return await todos.json();
}

// init function for app
(async function () {
  const todos = await fetchTodos();

  todos.forEach(({ id, title, complete }) => {
    const [tableRow, todoID, todoTitle, todoStatus] = [
      document.createElement("tr"),
      document.createElement("td"),
      document.createElement("td"),
      document.createElement("td"),
    ];

    todoID.innerHTML = id;
    todoTitle.innerHTML = title;
    todoStatus.innerHTML = complete ? "Complete" : "Pending";

    tableRow.append(...[todoID, todoTitle, todoStatus]);
    ui.tableBody.append(tableRow);
  });
})();
