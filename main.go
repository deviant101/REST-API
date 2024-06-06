package main

func main() {
	app := App{}
	app.Initialize(DBUser, DBPass, DBName)
	app.Run("localhost:8000")
}
