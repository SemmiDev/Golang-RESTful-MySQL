package student

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-mysql/config"
	models "golang-mysql/models"
	"log"
	"time"
)

const (
	table = "students"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Student, error) {
	var students []models.Student

	db, err := config.MySQL()
	if err != nil {
		log.Fatal("cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var student models.Student
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&student.ID,
			&student.Identifier,
			&student.Name,
			&student.Email,
			&student.Semester,
			&createdAt,
			&updatedAt); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		student.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		student.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}
	return students, nil
}

func Insert(ctx context.Context, student models.Student) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (identifier, name, email, semester, created_at, updated_at) values('%v','%v','%v','%v','%v', '%v')",
		table,
		student.Identifier,
		student.Name,
		student.Email,
		student.Semester,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func Update(ctx context.Context, studentUpdate models.Student) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set identifier = %s, name ='%s', email = '%s' ,semester = %d, updated_at = '%v' where id = %d",
		table,
		studentUpdate.Identifier,
		studentUpdate.Name,
		studentUpdate.Email,
		studentUpdate.Semester,
		time.Now().Format(layoutDateTime),
		studentUpdate.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id models.Student) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, id.ID)

	result, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := result.RowsAffected()

	if check == 0 {
		return errors.New("id not found ")
	}

	return nil
}
