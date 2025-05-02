import inquirer from "inquirer";
import {
  getAllTasks,
  createTask,
  getTaskById,
  updateTask,
  deleteTask,
  undoTaskUpdate,
} from "../api/tasks";
import { getAllUsers, createUser } from "../api/users";
import { checkHealth } from "../api/health";
import { checkStatus } from "../api/status";

export async function mainMenu() {
  while (true) {
    const { action } = await inquirer.prompt([
      {
        name: "action",
        type: "list",
        message: "Choose an action:",
        choices: [
          "📄 View Tasks",
          "🔍 View Task By ID",
          "➕ Create Task",
          "✏️ Edit Task",
          "❌ Delete Task",
          "↩️ Undo Task Update",
          "👤 View Users",
          "➕ Create User",
          "❤️ Check API Health",
          "🔧 Check Server Status",
          "🚪 Logout",
        ],
      },
    ]);

    switch (action) {
      case "📄 View Tasks":
        const tasks = await getAllTasks();
        console.table(tasks);
        break;
      case "🔍 View Task By ID":
        const { taskIdView } = await inquirer.prompt([
          { name: "taskIdView", message: "Enter Task ID:", type: "number" },
        ]);
        const task = await getTaskById(taskIdView);
        console.log(task);
        break;
      case "➕ Create Task":
        const newTaskInput = await inquirer.prompt([
          { name: "title", message: "Title:", type: "input" },
          { name: "description", message: "Description:", type: "input" },
          { name: "start", message: "Start Date (YYYY-MM-DD):", type: "input" },
          { name: "end", message: "End Date (YYYY-MM-DD):", type: "input" },
        ]);
        await createTask(
          newTaskInput.title,
          newTaskInput.description,
          newTaskInput.start,
          newTaskInput.end
        );
        console.log("✅ Task created!");
        break;
      case "✏️ Edit Task":
        const updateInput = await inquirer.prompt([
          { name: "id", message: "Task ID:", type: "number" },
          { name: "title", message: "New Title:", type: "input" },
          { name: "description", message: "New Description:", type: "input" },
          {
            name: "start",
            message: "New Start Date (YYYY-MM-DD):",
            type: "input",
          },
          { name: "end", message: "New End Date (YYYY-MM-DD):", type: "input" },
        ]);
        await updateTask(
          updateInput.id,
          updateInput.title,
          updateInput.description,
          updateInput.start,
          updateInput.end
        );
        console.log("✅ Task updated!");
        break;
      case "❌ Delete Task":
        const { idDelete } = await inquirer.prompt([
          { name: "idDelete", message: "Task ID to delete:", type: "number" },
        ]);
        await deleteTask(idDelete);
        console.log("🗑️ Task deleted.");
        break;
      case "↩️ Undo Task Update":
        const { idUndo } = await inquirer.prompt([
          { name: "idUndo", message: "Task ID to undo:", type: "number" },
        ]);
        await undoTaskUpdate(idUndo);
        console.log("↩️ Task update undone.");
        break;
      case "👤 View Users":
        const users = await getAllUsers();
        console.table(users);
        break;
      case "➕ Create User":
        const userInput = await inquirer.prompt([
          { name: "name", message: "Name:", type: "input" },
          { name: "email", message: "Email:", type: "input" },
          { name: "password", message: "Password:", type: "password" },
        ]);
        await createUser(userInput.name, userInput.email, userInput.password);
        console.log("✅ User created!");
        break;
      case "❤️ Check API Health":
        const health = await checkHealth();
        console.log("✅ API Health:", health);
        break;
      case "🔧 Check Server Status":
        const serverStatus = await checkStatus();
        console.log("✅ Server Status:", serverStatus);
        break;
      case "🚪 Logout":
        return;
    }
  }
}
