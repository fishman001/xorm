package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/positivelong/xorm"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need db path")
		return
	}

	orm, err := xorm.NewEngine("sqlite3", os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer orm.Close()
	orm.ShowSQL(true)

	tables, err := orm.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, table := range tables {
		fmt.Println(table.Name)
	}
}
