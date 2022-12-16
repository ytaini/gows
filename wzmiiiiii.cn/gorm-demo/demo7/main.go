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
	//err = InsertDemo1()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo1()

	//err = InsertDemo2()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo2()

	//err = InsertDemo3()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo3()

	//err = BatchInsertDemo1()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo4()

	//err = BatchInsertDemo2()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo5()

	//err = BatchInsertDemo3()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo6()

	//err = InsertByMapDemo1()
	//if err != nil {
	//	panic(err)
	//}
	//DryRunDemo7()

	//err = InsertByMapDemo2()
	//if err != nil {
	//	panic(err)
	//}

}
