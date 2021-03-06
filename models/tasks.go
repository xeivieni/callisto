package models

import (
	"github.com/graphql-go/graphql"
)

type Task struct {
	Id          int    `json:"id",db:"id"`
	ProjectId   int    `json:"projectId",db:"projectId"`
	Name        string `json:"name",db:"name"`
	Description string `json:"description",db:"description"`
	Type        int    `json:"type",db:"type"`
	StatusId    int    `json:"statusId",db:"statusId"`
	Deleted     int    `json:"deleted",db:"deleted"`
}

// define custom GraphQL ObjectType `TaskType` for our Golang struct `Task`
// Note that
// - the fields in our TaskType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var TaskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"projectId": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.Int,
		},
		"statusId": &graphql.Field{
			Type: graphql.Int,
		},
		"deleted": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func AllTasks() ([]*Task, error) {
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tsks := make([]*Task, 0)

	for rows.Next() {
		tsk := new(Task)
		err := rows.Scan(&tsk.Id, &tsk.ProjectId, &tsk.Name, &tsk.Description, &tsk.Type, &tsk.StatusId, &tsk.Deleted)
		if err != nil {
			return nil, err
		}
		tsks = append(tsks, tsk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tsks, nil
}

func TaskFromId(id int) (*Task, error) {
	row := db.QueryRow(`SELECT * FROM tasks WHERE id=$1`, id)

	tsk := new(Task)

	err := row.Scan(&tsk.Id, &tsk.ProjectId, &tsk.Name, &tsk.Description, &tsk.Type, &tsk.StatusId, &tsk.Deleted)
	if err != nil {
		return nil, err
	}
	return tsk, nil
}

func TasksForProject(id int) ([]*Task, error) {
	rows, err := db.Query(`SELECT * FROM tasks WHERE projectId=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tsks := make([]*Task, 0)

	for rows.Next() {
		tsk := new(Task)
		err := rows.Scan(&tsk.Id, &tsk.ProjectId, &tsk.Name, &tsk.Description, &tsk.Type, &tsk.StatusId, &tsk.Deleted)
		if err != nil {
			return nil, err
		}
		tsks = append(tsks, tsk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tsks, nil
}

func NewTask(tsk Task) (error) {
	stmt, err := db.Prepare("INSERT INTO tasks(projectId, name, description, type, statusId, deleted) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tsk.ProjectId, tsk.Name, tsk.Description, tsk.Type, tsk.StatusId, tsk.Deleted)
	if err != nil {
		return err
	}
	return nil
}

func MoveTask(id int, status int) (error) {
	// TODO: Add check on status
	stmt, err := db.Prepare("UPDATE tasks SET statusId=$1 WHERE id=$2")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, id)
	if err != nil {
		return err
	}
	return nil
}

func EditTask(tsk Task) (error) {
	// TODO: Add check on status
	stmt, err := db.Prepare("UPDATE tasks SET name=$1, description=$2, type=$3, statusId=$4 WHERE id=$5")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tsk.Name, tsk.Description, tsk.Type, tsk.StatusId, tsk.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int) (error) {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
