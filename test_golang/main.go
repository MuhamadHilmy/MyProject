package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	var resp string

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
	}

	if request["method"].(string) == "1000" {
		resp = mysql_getData()
	} else if request["method"].(string) == "1111" {
		query := "INSERT INTO students(name,age,gender,class) VALUES('" + request["nama"].(string) + "', '" + request["umur"].(string) + "', '" + request["jenisKelamin"].(string) + "', '" + request["kelas"].(string) + "')"
		resp = mysql_general(query)
	} else if request["method"].(string) == "1112" {
		query := "UPDATE students SET class = '" + request["kelas"].(string) + "' WHERE id = '" + request["id"].(string) + "'"
		resp = mysql_general(query)
	} else if request["method"].(string) == "1212" {
		query := "DELETE FROM students WHERE id = '" + request["id"].(string) + "'"
		resp = mysql_general(query)
	}

	w.Write([]byte(resp))
}

func main() {
	fmt.Println("SERVER STARTING ....")
	fmt.Println("SERVER STARTED, IP : 127.0.0.1 | PORT : 507")
	r := mux.NewRouter()

	r.HandleFunc("/test", handlerFunc)
	log.Fatal(http.ListenAndServe("127.0.0.1:507", r))
}
