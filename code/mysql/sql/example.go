package main

import "database/sql"

func GetNames(db *sql.DB) []string {

	sqlStr := "SELECT name FROM example_table" + " WHERE id > 0"
	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	l := []string{}
	var name *string
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&name)
			if err == nil {
				if name != nil {
					l = append(l, *name)
					name = nil
				}
			}
		}
	}

	if rows != nil {
		rows.Close()
	}

	return l
}

/*
	var numInSQL uint64
	row := db.QueryRow("SELECT MAX(theNumber) FROM transaction;")
	err = row.Scan(&numInSQL)
	if err != nil {
		panic(err.Error())
	}
*/
