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

	//err = InsertDemo()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo1()

	err = InsertConflictDemo1()
	if err != nil {
		panic(err)
	}
}
