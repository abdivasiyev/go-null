### go-null

Simple and lightweight library for working with sql null values

Installation

```bash
go get github.com/abdivasiyev/go-null
```

Usage

```go
package main

import (
	"github.com/abdivasiyev/go-null"
	"fmt"
	"database/sql"
)

type User struct {
	ID        int         `json:"id"`
	Firstname string      `json:"firstname"`
	Lastname  null.String `json:"lastname"`
	Age       null.Int    `json:"age"`
}

func main() {
	var u User
	// some stuff with sql

	if err := row.Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Age,
	); err != nil {
		panic(err)
	}

	fmt.Println(u.ID, u.Firstname, u.Lastname.Get(), u.Age.Get())
}
```