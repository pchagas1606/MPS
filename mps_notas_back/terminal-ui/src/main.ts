import { loginScreen } from "./ui/login";
import { mainMenu } from "./ui/mainMenu";

async function startApp() {
  console.clear();
  console.log("📋 Welcome to Task Manager Terminal UI");

  let success = false;
  while (!success) {
    success = await loginScreen();
  }

  await mainMenu();
  console.log("👋 Goodbye!");
}

startApp();
