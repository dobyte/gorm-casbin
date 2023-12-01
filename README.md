# gorm-casbin

gorm-based casbin component

Support MySQL, SQLite, PostgreSQL, Oracle, SQL Server Power By GoFrame ORM

## Use

Download and install

```shell
go get github.com/dobyte/gorm-casbin
```

Demo

```go
package main

import (
	"fmt"
	"log"
	"github.com/dobyte/gorm-casbin"
)

func main() {
	enforcer, err := casbin.NewEnforcer(&casbin.Options{
		Model:    "./model.conf",
		Debug:    false,
		Enable:   true,
		Autoload: true,
		Table:    "casbin_policy_test",
		Database: "root:123456@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local",
	})

	if err != nil {
		log.Fatalf("Casbin init failure:%s \n", err.Error())
	}

	// add a permission node for role
	ok, err := enforcer.AddPolicy("role_1", "node_1")

	if err != nil {
		log.Fatalf("Add policy exception:%s \n", err.Error())
	}

	if ok {
		log.Println("Add policy successful")
	} else {
		log.Println("Add policy failure")
	}
}
```

## Example

View demo [enforcer_test.go](enforcer_test.go)

## Model Demo

View demo [model.conf](model.conf)