package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := "112.74.172.81"
	port := "3306"
	user := "root"
	password := "123456"
	database := "mytestdb"

	hander := OpenSQL(user, password, host, port, database)
	hander.SetMaxOpenConns(100)
	hander.SetMaxIdleConns(50)
	err := hander.Ping()
	if err != nil {
		hander.Close()
		fmt.Println("Ping err=" + err.Error())
		time.Sleep(time.Second * 1)
		return
	}

	defer func() {
		if hander != nil {
			hander.Close()
		}
	}()

	names := GetNames(hander)
	fmt.Printf("%+v", names)
	fmt.Println()

	AllOperat = make(chan string, 20000)
	go SQLInsert(hander)
	saveToChan()

	now := time.Now()
	next := now.Add(time.Second * 5)
	ti := time.NewTimer(next.Sub(now))
	<-ti.C
}
