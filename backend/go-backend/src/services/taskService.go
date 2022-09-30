package services

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"main/src/models"
	//"log"
	"net/http"
)

var dbConnection *sql.DB

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	sqlStatement := `SELECT * FROM tasks`

	rows, err := dbConnection.Query(sqlStatement)

	if err == nil {
		var tasks []models.Task
		for rows.Next() {
			var task models.Task
			rows.Scan(&task.ID, &task.Label, &task.Done, &task.Date)
			tasks = append(tasks, task)
		}

		err := json.NewEncoder(writer).Encode(&tasks)
		if err != nil {
			return 
		}

	} else {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	//defer rows.Close()
	//defer dbConnection.Close()
}

func SetDB(database *sql.DB) {
	dbConnection = database
}


func CreateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var task models.Task
	var id string

	_ = json.NewDecoder(request.Body).Decode(&task)

	sqlStatement := `INSERT INTO tasks(id, label, done, date) VALUES($1, $2, $3, $4) RETURNING id`

	err := dbConnection.QueryRow(sqlStatement, task.ID, task.Label, task.Done, task.Date).Scan(&id)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	task.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(writer).Encode(&task)
}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id, _ := params["id"]

	var searchTask models.Task

	sqlStatement := `SELECT * FROM tasks WHERE id=$1`
	row := dbConnection.QueryRow(sqlStatement, id)

	row.Scan(&searchTask)
	json.NewEncoder(writer).Encode(&searchTask)

}


func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	var task models.Task
	_ = json.NewDecoder(request.Body).Decode(&task)
	oldID := params["id"]

	id := 0
	sqlStatement := `UPDATE tasks SET id=$1, label=$2, done=$3, date=$4 WHERE id=$5 RETURNING id`
	err := dbConnection.QueryRow(sqlStatement, task.ID, task.Label, task.Done, task.Date, oldID).Scan(&id)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
	log.Println("Updated record ID is:", id)

	json.NewEncoder(writer).Encode(&task)
}


func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id := params["id"]

	sqlStatement := `DELETE FROM tasks WHERE id=$1 RETURNING id`
	err := dbConnection.QueryRow(sqlStatement, id).Scan(&id)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
	log.Println("Deleted record ID is:", id)
	json.NewEncoder(writer).Encode(id)

}
