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
	Insert()
	//err = DeleteDemo1()
	//err = DeleteDemo2()
	err = BeforeDeleteDemo()
	if err != nil {
		panic(err)
	}
}

func Insert() {
	db.Create(GenerateUsers())
}
