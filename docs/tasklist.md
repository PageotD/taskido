# Taskido JSON Format

| Key           | Type     | Req |
|---------------|----------|-----|
| ID            | int      | Yes |
| UUID          | string   | Yes |
| Description   | string   | Yes |
| Projects      | []string | No  |
| Contexts      | []string | No  |
| Due           | string   | No  |
| Status        | string   | Yes |
| Priority      | int      | No  |
| CreatedAt     | string   | Yes |
| UpdatedAt     | string   | No  |

## Description:

* **`ID`**</br>A unique identifier for the task within the task list.

* **`UUID`**</br> A unique identifier for the task.

* **`Description`**</br> A brief description or title of the task.

* **`Projects`**</br> A list of projects that the task is associated with.

* **`Contexts`**</br> A list of contexts or tags present in the task description.

* **`Due`**</br> The due date of the task in the YYYY-MM-DD format.

* **`Status`**</br> A string indicating the status of the task (__pending__, __completed__, __archived__).

* **`Priority`**</br> An integer value representing the priority level of the task (_0_ for no priority, _1_ for low priority, _2_ for medium priority, _3_ for high priority).

* **`CreatedAt`**</br> The timestamp when the task was created, in the `YYYY-MM-DD HH:MM:SS` format.

* **`UpdatedAt`**</br> The timestamp when the task was last updated, in the `YYYY-MM-DD HH:MM:SS` format.