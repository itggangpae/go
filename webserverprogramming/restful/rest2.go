package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student //학생 목록을 저장하는 맵
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() //gorilla/mux를 만듭니다.
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	//-- 여기에 새로운 핸들러 등록 --//
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")

	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student) //임시 데이터 생성
	students[1] = Student{1, "adam", 53, 99}
	students[2] = Student{2, "jessica", 51, 87}
	students[3] = Student{3, "hunt", 13, 91}
	lastId = 3

	return mux
}

type Students []Student // Id로 정렬하는 인터페이스
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // ➎ 학생 목록을 Id로 정렬
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) //JSON 포맷으로 변경
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //id에 해당하는 학생이 없으면 에러
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student) // JSON 데이터 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++ // id를 증가시킨후 맵에 등록
	student.Id = lastId
	students[lastId] = student
	log.Println(students)
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // id에 해당하는 학생이 없으면 에러
		return
	}
	delete(students, id)
	log.Println(students)
	w.WriteHeader(http.StatusOK) // StatusOK 반환
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
