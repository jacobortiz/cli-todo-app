package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	commands := NewCommandFlags()
	commands.Execute(&todos)
	storage.Save(todos)
}
