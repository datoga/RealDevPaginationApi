package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type LogResponse struct {
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Fields `json:"results"`
}

type LogController struct {
	Port    int
	service *LogService
}

func NewLogController(port int, service *LogService) *LogController {
	return &LogController{
		Port:    port,
		service: service,
	}
}

func (controller LogController) Start() {
	http.HandleFunc("/logs/", controller.LogsHandler)
	http.HandleFunc("/logs", controller.LogsHandler)

	http.ListenAndServe(":"+strconv.Itoa(controller.Port), nil)
}

func (controller LogController) LogsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	log.Println(r.URL.Path)
	log.Printf("%v\n", r.URL.Query())
	log.Printf("%v\n", r.PostForm)

	cursors, ok := r.URL.Query()["cursor"]

	var models Models
	var previousCursor *string
	var nextCursor *string
	var err error

	if !ok {
		models, previousCursor, nextCursor, err = controller.service.Query("")
	} else {
		cursor := cursors[0]

		fmt.Println("Cursor is", cursor)

		models, previousCursor, nextCursor, err = controller.service.Query(cursor)
	}

	var path string

	if strings.HasSuffix(r.URL.Path, "/") {
		path = r.URL.Path[0 : len(r.URL.Path)-1]
	} else {
		path = r.URL.Path
	}

	var previousURL *string

	if previousCursor != nil {
		url := fmt.Sprintf("https://%s%s/?cursor=%s", r.Host, path, *previousCursor)
		previousURL = &url
	}

	var nextURL *string

	if nextCursor != nil {
		url := fmt.Sprintf("http://%s%s/?cursor=%s", r.Host, path, *nextCursor)
		nextURL = &url
	}

	fields := models.GetFields()

	logResponse := LogResponse{
		Previous: previousURL,
		Next:     nextURL,
		Results:  fields,
	}

	marshaledResponse, err := json.Marshal(logResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(string(marshaledResponse))

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledResponse)
	w.Write([]byte("\n"))
}
