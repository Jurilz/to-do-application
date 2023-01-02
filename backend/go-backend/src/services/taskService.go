package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	_ "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	_ "github.com/uptrace/bun/extra/bundebug"
	"log"
	"main/src/models"
	//"log"
	"net/http"
)

var dbConnection *bun.DB
var ctx = context.Background()

func SetDB(database *sql.DB) {
	//dbConnection = database
	dns := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	//dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//	os.Getenv("PGUSER"),
	//	os.Getenv("PGPASSWORD"),
	//	os.Getenv("PGHOST"),
	//	os.Getenv("PGPORT"),
	//	os.Getenv("PGDATABASE"),
	//)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))

	dbConnection = bun.NewDB(sqldb, pgdialect.New())

	dbConnection.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
}

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//sqlStatement := `SELECT * FROM tasks`
	//rows, err := dbConnection.Query(sqlStatement)
	var tasks []models.Task
	err := dbConnection.NewSelect().Model(&tasks).Scan(ctx)


	if err == nil {
/*		var tasks []models.Task
		for rows.Next() {
			var task models.Task
			rows.Scan(&task.ID, &task.Label, &task.Done, &task.Date)
			tasks = append(tasks, task)
		}*/

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




func CreateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var task models.Task
	var id int64

	_ = json.NewDecoder(request.Body).Decode(&task)

	//sqlStatement := `INSERT INTO tasks(id, label, done, date) VALUES($1, $2, $3, $4) RETURNING id`
	//err := dbConnection.QueryRow(sqlStatement, task.ID, task.Label, task.Done, task.Date).Scan(&id)

	result, err := dbConnection.NewInsert().Model(&task).Exec(ctx)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	task.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(writer).Encode(&result)
}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id, _ := params["id"]

	var searchTask models.Task

	//sqlStatement := `SELECT * FROM tasks WHERE id=$1`
	//row := dbConnection.QueryRow(sqlStatement, id)

	err := dbConnection.NewSelect().Model(&searchTask).Where("id = ?", id).Scan(ctx)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	//searchTask.Scan(&searchTask)
	json.NewEncoder(writer).Encode(&searchTask)

}


func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	var task models.Task
	_ = json.NewDecoder(request.Body).Decode(&task)
	oldID := params["id"]

	//id := 0
	//sqlStatement := `UPDATE tasks SET id=$1, label=$2, done=$3, date=$4 WHERE id=$5 RETURNING id`
	//err := dbConnection.QueryRow(sqlStatement, task.ID, task.Label, task.Done, task.Date, oldID).Scan(&id)
	res, err := dbConnection.NewUpdate().Model(&task).Where("id = ?", oldID).Exec(ctx)


	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
	//log.Println("Updated record ID is:", id)

	json.NewEncoder(writer).Encode(&res)
}


func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var task models.Task

	params := mux.Vars(request)
	id := params["id"]

	//sqlStatement := `DELETE FROM tasks WHERE id=$1 RETURNING id`
	//err := dbConnection.QueryRow(sqlStatement, id).Scan(&id)

	res, err := dbConnection.NewDelete().Model(&task).Where("id = ?", id).Exec(ctx)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
	log.Println("Deleted record ID is:", id)
	json.NewEncoder(writer).Encode(res)

}
