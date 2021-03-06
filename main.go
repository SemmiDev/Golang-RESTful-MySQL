package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-mysql/models"
	"golang-mysql/student"
	"golang-mysql/utils"
	"log"
	"net/http"
	"strconv"
)

func main(){

	http.HandleFunc("/students", GetStudents)
	http.HandleFunc("/students/create", PostStudent)
	http.HandleFunc("/students/update", UpdateStudent)
	http.HandleFunc("/students/delete", DeleteStudent)

	log.Println("SERVER STARTED ON PORT 9090")
	err := http.ListenAndServe(":9090", nil)

    if err != nil {
		log.Fatal(err)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Student

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		mhs.ID, _ = strconv.Atoi(id)

		if err := student.Delete(ctx, mhs); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Deleted",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Student

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := student.Update(ctx, mhs); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Updated",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var studentData models.Student

		if err := json.NewDecoder(r.Body).Decode(&studentData); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := student.Insert(ctx, studentData); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string {
			"status": "CREATED",
		}
		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		students, err := student.GetAll(ctx)
		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, students, http.StatusOK)
		return
	}
}
