package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// AllOperat 发送数据库查询语句
var AllOperat chan string

// MyData ...
type MyData struct {
	Name  string
	State int
}

func saveToChan() {
	myData := new(MyData)
	myData.Name = "Chan" + strconv.FormatInt(time.Now().Unix(), 10)
	myData.State = 1

	myData.Insert()
}

// Insert ...
func (thi *MyData) Insert() {

	var bufname, bufvalue bytes.Buffer

	bufname.WriteString("INSERT INTO example_table (")
	bufvalue.WriteString(" values(")
	if thi.Name != "" {
		bufname.WriteString("name")
		bufvalue.WriteString("'" + thi.Name + "'")
	}
	if thi.State != 0 {
		bufname.WriteString(", state")
		bufvalue.WriteString("," + strconv.Itoa(thi.State))
	}
	bufname.WriteString(")")
	bufvalue.WriteString(")")
	bufname.WriteString(bufvalue.String())

	fmt.Println()
	fmt.Println(bufname.String())
	fmt.Println()

	AllOperat <- bufname.String()
}

// SQLInsert ...
func SQLInsert(db *sql.DB) {
	for query := range AllOperat {
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
