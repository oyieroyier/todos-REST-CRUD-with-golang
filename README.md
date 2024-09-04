# CRUD OPERATIONS USING GO

This is a quick practice on working with Go (Golang) to build a REST API.

## System Requirements

- Go 1.22+
- Postman/Thunder Client
- Operating System (Windows 10+, MacOS, Linux, etc.)
- A text editor capable of running Go (Visual Studio Code, Vim, Nano, Emacs, Atom, Sublime Text, etc.)
- RAM >= 4GB
- Disk space >= 1GB

## Installation

To use this repo, follow these steps:

### Alternative One

1.  Open the terminal/CLI on your computer.

2.  Clone the repository by running the following command:

        git clone https://github.com/oyieroyier/todos-REST-CRUD-with-golang.git

3.  Change directory to the repo folder:

        cd todos-REST-CRUD-with-golang

4.  Open it in your Code editor of choice. If you use VS Code, run the command:

        code .

### Alternative Two

- In the top right corner of this page there is a button labelled `Fork`.

- Click on that button to create a copy of the repository to your own account.

- Follow the process described in `Alternative One` above.

- Remember to replace your username when cloning.

      git clone https://github.com/your-github-username-here/todos-REST-CRUD-with-golang.git

## Running the application locally

- Open the integrated terminal in your code editor and run the application:

      go run main.go

- This will open the development server in your terminal.

- Open Postman/Thunder Client and paste the below URL:

      http://localhost:8080/

## `GET` all todos.

    /todos

### Response:
```json
[
    {
        "id": "1",
        "title": "Learn some Go",
        "completed": true
    },
    {
        "id": "2",
        "title": "Build a REST API in Go",
        "completed": true
    },
    {
        "id": "3",
        "title": "Deploy the REST API",
        "completed": false
    }
]
```

## `GET` single todo

    /todos/:id

### Response:
```json
    {
        "id": "1",
        "title": "Learn some Go",
        "completed": true
    }
```
### Error: (nonexistent todo)
```json
    {
        "message": "selected todo not found"
    }
```

## `POST` a new todo.

    /todos

### Response:
```json
    {
      "id": "4",
      "title": "Find out how to implement hot reload in Go",
      "completed": false
    }
```


## Authors

[Bob Oyier](https://github.com/oyieroyier/)
