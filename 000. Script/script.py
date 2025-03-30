import os
import shutil
import subprocess

main_file = "C:/Users/Resurrected/Learn_Go/main.go"
repo_folder = "C:/Users/Resurrected/my/"

folder_count = len(next(os.walk(repo_folder))[1]) - 1

folder_name = input("Enter folder name: ")
new_folder = f"{repo_folder}/{folder_count:03}. {folder_name}"
os.makedirs(new_folder, exist_ok=True)
destination_file = f"{new_folder}/main.go"
shutil.copyfile(main_file, destination_file)

try:
    subprocess.run(["git", "add", "-A"], check=True, cwd=repo_folder)
    subprocess.run(["git", "commit", "-m", folder_name], check=True, cwd=repo_folder)
    subprocess.run(["git", "push", "-u", "origin", "main"], check=True, cwd=repo_folder)

    print("Changes pushed successfully!")
except subprocess.CalledProcessError as e:
    print(f"Error: {e}")

go_code = """package main

import "fmt"

func main() {

}
"""

with open(main_file, "w") as f:
    f.write(go_code)