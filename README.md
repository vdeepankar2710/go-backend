# go-backend

Todo App in Golang

## Runing the project.

1. Setup golang in your system by downloading golang and setting up environment variables if neccessary.
2. Open the project in VSCode and and run `go mod tidy` for all the packages to be installed for running the project.
3. Run the project by `go run main.go`

## The project was asked to be made using SkyllaDB as the database but I used MongoDB as I am using Windows and SkyllaDB is not available for Windows platform

## Testing the APIs:

1. Endpoint: `{your local machine}/todos/create`- Creates a todo entry.

### POST Request:

`{
    "title":"7 todo",
    "description":"todosfsaf 8",
    "user_id":1,
    "status":"Not Completed"
}`

### Response :

`{
"message":"Todo created successfully",
"todo":{
"id":"66693941a8fd0fba1cac5804",
"user_id":1,
"title":"7 todo",
"description":"todosfsaf 8",
"status":"Not Completed","created_at":"2024-06-12T11:29:29.0014676+05:30","updated_at":"2024-06-12T11:29:29.0014676+05:30"}}

`

2. Endpoint: `{your local machine}/todos/get/1/6/ASC`- Get todo entries on page 1 with 6 entries per page sorted in Ascending (ASC) order OR for descending use (DESC). The above query params are for pagination.

### GET Request:

### Response :

`[{
"id":"6668702b151ec76c39a62c6a",
"user_id":1,
"title":"First todo",
"description":"todo desc1",
"status":"Not Completed","created_at":"2024-06-11T15:41:31.084Z","updated_at":"2024-06-11T15:41:31.084Z"},
{
"id":"66687086151ec76c39a62c6b",
"user_id":1,
"title":"second todo",
"description":"todo desc2",
"status":"Not Completed","created_at":"2024-06-11T15:43:02.884Z","updated_at":"2024-06-11T15:43:02.884Z"
},
{
"id":"666870a1f5bcbe36274c8874",
"user_id":1,
"title":"second todo",
"description":"todo desc2",
"status":"Not Completed","created_at":"2024-06-11T15:43:29.246Z","updated_at":"2024-06-11T15:43:29.246Z"
},
{
"id":"66687155f5bcbe36274c8875",
"user_id":1,
"title":"third todo",
"description":"todo desc3",
"status":"Not Completed","created_at":"2024-06-11T15:46:29.906Z","updated_at":"2024-06-11T15:46:29.906Z"
},
{
"id":"666872b458b1ea308b04411b",
"user_id":1,
"title":"4 todo",
"description":"todo desc4",
"status":"Not Completed","created_at":"2024-06-11T15:52:20.372Z","updated_at":"2024-06-11T15:52:20.372Z"
},
{
"id":"6668736b933a48afcbeca43c",
"user_id":1,
"title":"5 todo",
"description":"todo desc5",
"status":"Not Completed","created_at":"2024-06-11T15:55:23.606Z","updated_at":"2024-06-11T15:55:23.606Z"}]`

3. Endpoint: `{your local machine}/todos/get/6668736b933a48afcbeca43c`- Get a todo by id.

### GET Request:

### Response :

`{
"message":"Todos retrieved successfully",
"todo":{"id":"6668702b151ec76c39a62c6a",
"user_id":1,
"title":"First todo",
"description":"todo desc1",
"status":"Not Completed","created_at":"2024-06-11T15:41:31.084Z","updated_at":"2024-06-11T15:41:31.084Z"}}

`

4. Endpoint: `{your local machine}/todos/update/6668736b933a48afcbeca43c`- Update a todo by id.

### PUT Request:

`{
    "title":"New title for this todo"
}`

### Response :

`
{
"id":"6668736b933a48afcbeca43c",
"user_id":1,
"title":"New title for this todo","description":"todo desc5",
"status":"Not Completed","created_at":"2024-06-11T15:55:23.606Z","updated_at":"2024-06-12T11:14:23.4517579+05:30"}

`

5. Endpoint: `{your local machine}/todos/delete/6668736b933a48afcbeca43c`- Delete a todo by id.

### DELETE Request:

### Response :

`
`
