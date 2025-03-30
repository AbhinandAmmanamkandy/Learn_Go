import os
import shutil

main_file = r"C:\Users\Resurrected\Learn_Go\main.go"
repo_folder = r"C:\Users\Resurrected\repository\\"

folder_count = sum(len(dirs) for _, dirs, _ in os.walk(repo_folder)) + 1

folder_name = input("Enter folder name: ")
new_folder = f"{repo_folder}{folder_count:03}. {folder_name}"
os.makedirs(new_folder, exist_ok=True)
shutil.copyfile(main_file, new_folder)
