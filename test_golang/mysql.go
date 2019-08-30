package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlHost = "127.0.0.1"
var mysqlPort = "3306"
var mysqlUser = "root"
var mysqlPass = ""
var mysqlDb = "school"

func mysql_getData() string {
	var id, name, age, gender, class string
	resp := map[string]interface{}{}
	var mapData []map[string]interface{}

	con, err := sql.Open("mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlHost+":"+mysqlPort+")/"+mysqlDb)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()

	rows, err := con.Query("SELECT * FROM students")
	if err != nil {
		resp["status"] = "14"
		resp["description"] = "DATA NOT FOUND"
	} else {
		for rows.Next() {
			err = rows.Scan(&id, &name, &age, &gender, &class)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				data := map[string]interface{}{
					"id":     id,
					"name":   name,
					"age":    age,
					"gender": gender,
					"class":  class,
				}

				mapData = append(mapData, data)
			}
		}
		resp["status"] = "00"
		resp["description"] = "GET DATA SUCCESS"
		resp["data"] = mapData
	}
	datas, _ := json.Marshal(mapData)
	resps, _ := json.Marshal(resp)
	fmt.Println("DATA DB : " + string(datas))
	return string(resps)
}

func mysql_general(query string) string {
	resp := map[string]interface{}{}

	con, err := sql.Open("mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlHost+":"+mysqlPort+")/"+mysqlDb)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()

	stmt, err := con.Prepare(query)
	if err != nil {
		resp["status"] = "14"
		resp["description"] = "INSERT/UPDATE/DELETE DATA, FAILED"
	} else {
		_, err := stmt.Exec()
		if err != nil {
			resp["status"] = "14"
			resp["description"] = "INSERT/UPDATE/DELETE DATA, FAILED"
		} else {
			resp["status"] = "00"
			resp["description"] = "INSERT/UPDATE/DELETE DATA, SUCCESS"
		}
	}
	resps, _ := json.Marshal(resp)
	fmt.Println(resp["description"].(string))
	return string(resps)
}
