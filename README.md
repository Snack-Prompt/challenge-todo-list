# TODO List Challenge

## Test Objective:
Create a basic REST API in Golang using the Fiber framework and PostgreSQL as the database. The API will manage a "to-do list" system.

## Requirements:

### 1. Setup and Configuration:
- Install Fiber as the framework.
- Set up PostgreSQL with a table for tasks.

### 2. API Functionalities:
- **Create Task**: An endpoint to create a new task.
  - Method: `POST /tasks`
  - Request body: `{"title": "Learn Golang", "description": "Study Goroutines", "status": "pending"}`

- **List Tasks**: An endpoint to list all tasks.
  - Method: `GET /tasks`
  - Response: JSON containing all tasks.

- **Update Task**: An endpoint to update the status of a task.
  - Method: `PUT /tasks/:id`
  - Request body: `{"status": "completed"}`

- **Delete Task**: An endpoint to delete a task.
  - Method: `DELETE /tasks/:id`

### 3. Database:
- Create a `tasks` table with the following fields:
  - `id` (UUID)
  - `title` (string)
  - `description` (string)
  - `status` (string) - possible values: `pending`, `completed`
  - `created_at` (timestamp)
  - `updated_at` (timestamp)

### 4. Validation:
- Ensure required fields are present when creating or updating a task.
- Return appropriate error responses (400, 404) when something goes wrong (e.g., trying to update or delete a non-existing task).

### 5. Extras (Optional):
- Add pagination to the task listing endpoint.
- Add support for filtering tasks by status (pending or completed).
- Add Unit Tests.
- Add Dockerfile.
- Create a Postman Collection with all endpoints.

## How to Submit:
1. Fork this repository.
2. Clone the forked repository to your machine:
   ```bash
   git clone https://github.com/Snack-Prompt/challenge-todo-list.git
3. Open a Pull Request to this repostory.
