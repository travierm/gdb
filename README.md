# gdb
A basic database class written in Go

## Basic Operations
```go
db := MapDatabase{Name: "Cache"}

db.Add("int", 8)
db.Add("string", "This is a String")

myInt := db.Get("int")
myString := db.Get("string")
```

