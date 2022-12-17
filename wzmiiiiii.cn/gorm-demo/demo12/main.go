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
	//err = SelectDemo1()
	//err = SelectDemo2()
	//err = SelectDemo3()
	err = SelectDemo4()
	if err != nil {
		panic(err)
	}
}
