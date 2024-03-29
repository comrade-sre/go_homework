package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Handler struct {
}
type Employee struct {
	Name   string  `json:"name" xml:"name"`
	Age    int     `json:"age" xml:"age"`
	Salary float32 `json:"salary" xml:"salary"`
}
type UploadHandler struct {
	HostAddr string
	UploadDir string
}
type GetHandler struct {
	UploadDir string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.FormValue("name")
		fmt.Fprintf(w, "Parsed param name with value: %s", name)
	case http.MethodPost:
		defer r.Body.Close()
		contentType := r.Header.Get("Content-Type")
		var employee Employee
		switch contentType {
		case "application/json":
			err := json.NewDecoder(r.Body).Decode(&employee)
			if err != nil {
				http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
				return
			}

		case "application/xml":
			err := xml.NewDecoder(r.Body).Decode(&employee)
			if err != nil {
				http.Error(w, "Unable to unmarshal XML", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Unknown content type", http.StatusBadRequest)
			return

		}
		fmt.Fprintf(w, "Got a new employee\nName: %s\nAge: %d\nSalary: %0.2f\n",
			employee.Name,
			employee.Age,
			employee.Salary,
		)

	}
}
func (h *GetHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	files, err := ioutil.ReadDir("./"+ h.UploadDir)
	if err != nil {
		fmt.Fprintf(w, "unable to read dir: %v", err )
		return
	}
	for _, file := range files {
		fileSlice := strings.Split(file.Name(), ".")
		if len(fileSlice) > 1 {
			extension := fileSlice[len(fileSlice)-1]
			fmt.Fprintf(w, "name: %s\t type: %s\t size: %d\n",fileSlice[0], extension, file.Size())
		} else {
			fmt.Fprintf(w, "name: %s\t type: no type\t size: %d\n",file.Name(), file.Size())
		}
	}
}
func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	filepath := h.UploadDir + "/" + header.Filename
	err = ioutil.WriteFile(filepath, data, 0755)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	h.HostAddr = "http://" + r.Host + ":8080"
	fileLink :=  h.HostAddr+ "/" + header.Filename
	fmt.Fprintln(w, fileLink)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	mux := http.NewServeMux()
	handler := &Handler{}
	uploadHandler := &UploadHandler{
		UploadDir: "upload",
	}
	mux.Handle("/upload", uploadHandler)
	mux.Handle("/", handler)
	srv := &http.Server{
		Addr:         ":80",
		Handler: mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
		wg.Done()
	}()
	gethandler := &GetHandler{
		UploadDir: uploadHandler.UploadDir,
	}
	dirToServe := http.Dir(uploadHandler.UploadDir)
	filemux := http.NewServeMux()
	filemux.Handle("/", http.FileServer(dirToServe))
	filemux.Handle("/ls", gethandler)
	fs := &http.Server{
		Addr: ":8080",
		Handler: filemux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		log.Fatal(fs.ListenAndServe())
		wg.Done()
	}()
	wg.Wait()
}

