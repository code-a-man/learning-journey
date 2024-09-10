# 001 Task Tracker

A simple task tracker that uses CLI to manage tasks. Based on the [Task Tracker](https://roadmap.sh/projects/task-tracker) project.

## Features

- Add tasks
- List tasks
  - List all tasks that are done, in progress, or not done
- Update tasks
  - Mark tasks as done, in progress, or not done
  - Update task description
- Delete tasks
- Save tasks to a JSON file and load them back

## Usage

```bash
# Adding a new task
task-tracker add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-tracker update 1 "Buy groceries and cook dinner"
task-tracker delete 1

# Marking a task as in progress or done
task-tracker mark 1 in-progress
task-tracker mark 1 done
task-tracker mark 1 todo

# Listing all tasks
task-tracker list

# Listing tasks by status
task-tracker list done
task-tracker list todo
task-tracker list in-progress
```
