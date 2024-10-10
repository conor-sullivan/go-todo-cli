# Go Todo App for cli

## Description

Go Todo App for cli is a simple command-line application written in Go that allows you to manage your tasks effectively. With this app, you can add, delete, edit, toggle, and list your todos right from your terminal.

## Features

- **Add a new todo**
- **Delete a todo by index**
- **Edit a todo by index**
- **Toggle the completion status of a todo**
- **List all todos**

## Installation

To use this app, you need to have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

1. Clone the repository:

   ```bash
   git clone https://github.com/conor-sullivan/go-todo-cli
   ```

2. Change into the project directory:

   ```bash
   cd go-todo-cli
   ```

3. Build the application:

   ```bash
   go build
   ```

## Usage

The app uses command-line flags to execute commands. Here’s how to use it:

### Commands

- **Add a new todo**

  ```bash
  ./go-todo-cli -add "Your Todo Title"
  ```

- **Delete a todo by index**

  ```bash
  ./go-todo-cli -del <index>
  ```

- **Edit a todo**

  ```bash
  ./go-todo-cli -edit "<index>:new_title"
  ```

- **Toggle a todo as completed or not completed**

  ```bash
  ./go-todo-cli -toggle <index>
  ```

- **List all todos**

  ```bash
  ./go-todo-cli -list
  ```

### Example Usage

1. Add a todo:

   ```bash
   ./go-todo-cli -add "Buy groceries"
   ```

2. List todos:

   ```bash
   ./go-todo-cli -list
   ```

3. Edit a todo (assuming the index of "Buy groceries" is 0):

   ```bash
   ./go-todo-cli -edit "0:Buy groceries and cook dinner"
   ```

4. Toggle the completion status of the todo with the index of 0:

   ```bash
   ./go-todo-cli -toggle 0
   ```

5. Delete a todo with an index of 0:

   ```bash
   ./go-todo-cli -del 0
   ```

## Contributing

Feel free to submit issues and pull requests! Contributions are welcome. 

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Commit your changes and push to your branch.
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

