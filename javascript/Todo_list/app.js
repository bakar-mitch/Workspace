const addButton = document.getElementById("addButton");
const taskInput = document.getElementById("task");
const priorityInput = document.getElementById("priority");
const dueDateInput = document.getElementById("dueDate");
const startDateInput = document.getElementById("startDate");
const taskList = document.getElementById("taskList");
let draggedTaskIndex;

// Load tasks from local storage on page load
document.addEventListener("DOMContentLoaded", () => {
  const storedTasks = JSON.parse(localStorage.getItem("tasks")) || [];
  storedTasks.forEach(task => {
    addTask(task.text, task.priority, task.dueDate, task.startDate, task.completed);
  });
});

addButton.addEventListener("click", () => {
  const taskText = taskInput.value.trim();
  const priority = priorityInput.value;
  const dueDate = dueDateInput.value;
  const startDate = startDateInput.value;
  const completed = false; // New task is not completed

  if (taskText !== "") {
    addTask(taskText, priority, dueDate, startDate, completed);
    clearInputFields();
    saveTasksToLocalStorage();
  }
});

function addTask(taskText, priority, dueDate, startDate, completed) {
  const taskRow = document.createElement("tr");
  taskRow.dataset.priority = priority;
  taskRow.dataset.dueDate = dueDate;
  taskRow.dataset.startDate = startDate;
  taskRow.dataset.completed = completed;
  taskRow.innerHTML = `
    <td>${taskText}</td>
    <td>${priority}</td>
    <td>${dueDate}</td>
    <td>${startDate}</td>
    <td>
      ${completed ? "Completed" : "<button class='btn btn-success btn-sm complete-button'>Mark as Completed</button>"}
    </td>
    <td>
      <button class="btn btn-success btn-sm edit-button">Edit</button>
      <button class="btn btn-danger btn-sm delete-button">Delete</button>
    </td>
  `;

  const deleteButton = taskRow.querySelector(".delete-button");
  deleteButton.addEventListener("click", () => {
    taskList.removeChild(taskRow);
    saveTasksToLocalStorage();
  });

  const editButton = taskRow.querySelector(".edit-button");
  editButton.addEventListener("click", () => {
    const newText = prompt("Edit task:", taskText);
    if (newText !== null && newText.trim() !== "") {
      taskRow.querySelector("td").textContent = newText;
      saveTasksToLocalStorage();
    }
  });

  const completeButton = taskRow.querySelector(".complete-button");
  if (!completed) {
    completeButton.addEventListener("click", () => {
      taskRow.dataset.completed = true;
      completeButton.textContent = "Completed";
      saveTasksToLocalStorage();
    });
  } else {
    completeButton.textContent = "Completed";
  }

  taskRow.setAttribute("draggable", true);

  taskRow.addEventListener("dragstart", event => {
    draggedTaskIndex = Array.from(taskList.children).indexOf(taskRow);
    event.dataTransfer.setData("text/plain", "");
  });

  taskRow.addEventListener("dragover", event => {
    event.preventDefault();
    const currentTaskIndex = Array.from(taskList.children).indexOf(taskRow);

    if (currentTaskIndex !== draggedTaskIndex) {
      taskList.insertBefore(
        draggedTaskIndex < currentTaskIndex ? taskRow : taskRow.nextElementSibling,
        taskList.children[draggedTaskIndex]
      );
      saveTasksToLocalStorage();
      draggedTaskIndex = currentTaskIndex;
    }
  });

  taskRow.addEventListener("drop", event => {
    event.preventDefault();
  });

  taskList.appendChild(taskRow);
  saveTasksToLocalStorage();
}

function saveTasksToLocalStorage() {
  const tasks = [];
  taskList.querySelectorAll("tr").forEach(taskRow => {
    tasks.push({
      text: taskRow.querySelector("td").textContent,
      priority: taskRow.dataset.priority,
      dueDate: taskRow.dataset.dueDate,
      startDate: taskRow.dataset.startDate,
      completed: taskRow.dataset.completed === "true"
    });
  });
  localStorage.setItem("tasks", JSON.stringify(tasks));
}

function clearInputFields() {
  taskInput.value = "";
  priorityInput.value = "low";
  dueDateInput.value = "";
  startDateInput.value = "";
}

