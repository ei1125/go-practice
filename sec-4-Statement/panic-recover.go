package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unaable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
