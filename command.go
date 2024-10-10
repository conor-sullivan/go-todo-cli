package main

import (
  "flag"
  "strings"
  "os"
  "fmt"
  "strconv"
)

type CmdFlags struct {
  Add string
  Del int
  Edit string
  Toggle int
  List bool
}

func NewCmdFlags() *CmdFlags {
  cf := CmdFlags{}

 flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title.")
 flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index.")
 flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
 flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo as completed or not completed.")
 flag.BoolVar(&cf.List, "list", false, "List all todos.")

 flag.Parse()

 return &cf
}

func (cf *CmdFlags) Execute (todos *Todos) {
  switch {
    default: 
      todos.print()
    case cf.List:
      todos.print()
    case cf.Add != "":
      todos.add(cf.Add)
      todos.print()
    case cf.Toggle != -1:
      todos.toggle(cf.Toggle)
      todos.print()
    case cf.Del != -1:
      todos.delete(cf.Del)
      todos.print()
    case cf.Edit != "":
      parts := strings.SplitN(cf.Edit, ":", 2)
      if len(parts) != 2 {
        fmt.Println("Error, invalid format for edit. Please use id:new_title.")
        os.Exit(1)
      }

      index, err := strconv.Atoi(parts[0])

      if err != nil {
        fmt.Println("Error: invalid index for edit")
        os.Exit(1)
      }

      todos.edit(index, parts[1])
      todos.print()
  }
}
