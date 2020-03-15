package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicOperations(t *testing.T) {

	db := MapDatabase{Name: "Cache"}
	m := make(map[int]string)

	m[1] = "Amy"
	m[2] = "Bob"
	m[3] = "Chis"

	db.Add("int", 2)
	db.Add("users", m)
	db.Add("string", "TestingString")

	users := db.Get("users").(map[int]string)

	assert.Equal(t, db.Get("int"), 2, "DB.Get should return the int 2")
	assert.Equal(t, users[1], "Amy", "DB.Get should return the string Amy")
	assert.Equal(t, db.Get("string"), "TestingString", "DB.Get should return the string TestingString")

}

func BenchmarkBasicOperations(b *testing.B) {
	db := MapDatabase{Name: "Cache"}

	data := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < b.N; i++ {
		db.Add(strconv.Itoa(i), data)
	}

	for i := 0; i < b.N; i++ {
		db.Get(strconv.Itoa(i))
	}
}
