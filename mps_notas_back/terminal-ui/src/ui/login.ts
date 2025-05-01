import inquirer from "inquirer";
import { login } from "../api/auth";

export async function loginScreen() {
  const { email, password } = await inquirer.prompt([
    { name: "email", message: "Email:", type: "input" },
    { name: "password", message: "Password:", type: "password" },
  ]);

  try {
    const result = await login(email, password);
    console.log(`✅ Login successful! Welcome ${email}`);
    return true;
  } catch {
    console.log("❌ Login failed. Try again.");
    return false;
  }
}
