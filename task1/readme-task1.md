# OOP Task Manager in Golang

This is a simple task manager implemented in Golang using Object-Oriented Programming (OOP) principles. It allows executing different types of tasks, monitoring their execution, and retrying failed tasks.

## Features

- Implements the `Task` interface.
- Three task types: `EmailTask`, `SMSTask`, and `ReportTask`.
- Logs task execution status (success or failure).
- Retries failed tasks.

## How It Works

1.  The program creates a `TaskManager`.
2.  It adds different types of tasks to the manager.
3.  The `TaskManager` executes all tasks one by one.
4.  If a task fails, it is stored for a retry.
5.  After initial execution, failed tasks are retried once.

## Installation & Usage

1.  Ensure you have Go installed.
2.  Clone this repository or copy the `main.go` file.
3.  Run the program using:

    ```bash
    go run main.go
    ```

## Code Explanation

### `Task` Interface

Defines two methods:

-   `Execute() error` - Runs the task.
-   `Info() string` - Returns task information.

### `BaseTask` Struct

A base struct containing common task attributes:

-   `ID` - Unique identifier.
-   `CreatedAt` - Timestamp of creation.
-   `Description` - Task description.

### Task Types

Each task type (`EmailTask`, `SMSTask`, `ReportTask`) embeds `BaseTask` and implements the `Task` interface.

### `TaskManager` Struct

Manages the task execution:

-   `AddTask(Task)` - Adds a new task.
-   `ExecuteAll()` - Executes all tasks sequentially.
-   `RetryFailedTasks()` - Retries failed tasks.

## Example Output

```bash
Executing EmailTask: 1
[SUCCESS] EmailTask completed in 2s
Executing SMSTask: 2
[ERROR] SMSTask failed: SMSTask 2 failed
Executing ReportTask: 3
[SUCCESS] ReportTask completed in 3s
Retrying failed tasks...
Executing SMSTask: 2
[SUCCESS] SMSTask completed in 1s
```

## Notes

-   Task execution time is random (1-3 seconds).
-   Tasks fail with a 20% probability (random simulation).