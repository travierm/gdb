package main

import "fmt"

func echo(s interface{}) {
	fmt.Println(s)
}

func main() {
	echo("Booted up Testing Environment")

	db := MapDatabase{Name: "Cache"}
	m := make(map[int]string)

	m[1] = "Amy"
	m[2] = "Bob"
	m[3] = "Chis"

	db.Add("users", m)
	users := db.Get("users").(map[int]string)

	echo(users[1])
}
