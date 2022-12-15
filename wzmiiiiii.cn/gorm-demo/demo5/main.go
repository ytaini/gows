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
	//InsertDemo1()
	//InsertDemo2()
	//UpdateCreatedAtDemo()
	UpdateUpdatedAtDemo()
}
