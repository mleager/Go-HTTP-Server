package main

func main() {
	app := App{}
	app.Initialize(DBUser, DBPassword, DBName)
	app.Run("localhost:8000")
}

//           Steps
// -------------------------
// 1. "go build"
//
// 2. ./<root directory>
//
// 3. localhost:8000/products
