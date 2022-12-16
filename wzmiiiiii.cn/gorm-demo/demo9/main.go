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
	//Insert()
	//err = SaveDemo()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo1()
	//DryRunDemo2()

	//err = UpdateDemo1()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo3()

	//err = UpdateDemo2()
	//if err != nil {
	//	panic(err)
	//}

	//err = UpdateDemo3()
	//if err != nil {
	//	panic(err)
	//}

	err = UpdateDemo4()
	if err != nil {
		panic(err)
	}
}

func Insert() {
	db.Create(GenerateUsers())
}
