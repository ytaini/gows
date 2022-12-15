package main

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	err = CreateTableDemo()
	if err != nil {
		panic(err)
	}
	err = InsertDemo()
	if err != nil {
		panic(err)
	}
}
