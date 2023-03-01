package server

import (
	"ddd/domain"
	service "ddd/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpServer struct {
	userSvc service.UserService
	taskSvc service.TaskService
}

func NewHttpServer(userSvc service.UserService, taskSvc service.TaskService) *httpServer {
	return &httpServer{
		userSvc: userSvc,
		taskSvc: taskSvc,
	}
}

func (s *httpServer) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/ping", s.handlePing).Methods(http.MethodGet)

	// tasks
	r.HandleFunc("/tasks", s.handleCreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{taskID}", s.handleGetOneTask).Methods(http.MethodGet)

	// users
	r.HandleFunc("/users", s.handleCreateUser).Methods(http.MethodPost)

	return http.ListenAndServe(":8080", r)
}

func (s *httpServer) handlePing(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("pong")
}

func (s *httpServer) handleCreateTask(w http.ResponseWriter, req *http.Request) {
	var input domain.Task

	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := s.taskSvc.CreateTask(input.UserID, input.Title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func (s *httpServer) handleGetOneTask(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	taskIDStr := vars["taskID"]

	taskID, err := strconv.Atoi(taskIDStr)

	log.Printf("Finding task: %v", taskID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := s.taskSvc.GetTaskByID(taskID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func (s *httpServer) handleCreateUser(w http.ResponseWriter, req *http.Request) {
}
