const ui = {
  inputElement: document.getElementById("title"),
  addButtonElement: document.getElementById("add-btn"),
  tableBody: document.querySelector("tbody"),
  formElement: document.querySelector("form"),
};

const serverURL = "http://localhost:5000/todos";

// function to fetch todos
async function fetchTodos() {
  const todos = await fetch(serverURL);

  return await todos.json();
}

// function to create todo
async function createTodo(title) {
  try {
    const response = await fetch(serverURL, {
      method: "POST",
      body: JSON.stringify({ title }),
    });

    const todo = await response.json();

    createTodoUi(todo);
  } catch {
    console.error("Something went wrong!");
  }
}

// init function for app
(async function () {
  const todos = await fetchTodos();

  todos.forEach((todo) => createTodoUi(todo));
})();

ui.formElement.addEventListener("submit", (e) => {
  e.preventDefault();
  createTodo(ui.inputElement.value);
});

function createTodoUi({ id, title, complete }) {
  let [
    tableRow,
    todoID,
    todoTitle,
    todoStatus,
    todoActions,
    markAsCompleteBtn,
    deleteBtn,
  ] = [
    document.createElement("tr"),
    document.createElement("td"),
    document.createElement("td"),
    document.createElement("td"),
    document.createElement("td"),
    document.createElement("button"),
    document.createElement("button"),
  ];

  // setting the data
  todoID.innerHTML = id;
  todoTitle.innerHTML = title;
  todoStatus.innerHTML = complete ? "Complete" : "Pending";

  // todo actions
  markAsCompleteBtn.innerHTML = "✅";
  deleteBtn.innerHTML = "❌";

  // adding ids for styling
  markAsCompleteBtn.id = "check-btn";
  deleteBtn.id = "del-btn";

  // adding event listener to each btn
  markAsCompleteBtn.addEventListener("click", () => {
    todoTitle.style.textDecoration = "line-through";
    todoStatus.innerHTML = "Complete";
  });

  deleteBtn.addEventListener("click", () => {
    // set all the feilds to null when item is deleted
    [todoActions, todoTitle, todoStatus, todoID] = [null, null, null, null];
    console.log([todoActions, todoTitle, todoStatus, todoID]);
  });

  // appending the btns to the actions row
  todoActions.append(...[markAsCompleteBtn, deleteBtn]);

  tableRow.append(...[todoID, todoTitle, todoStatus, todoActions]);
  ui.tableBody.append(tableRow);
}
