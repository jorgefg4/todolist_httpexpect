package server

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"

	//"ToDoList/assets"
	//"ToDoList/task"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/jorgefg4/todolist/pkg/task"
	"github.com/jorgefg4/todolist/web/assets"
)

// Templates
var navigationBarHTML string
var homepageTpl *template.Template
var secondViewTpl *template.Template
var thirdViewTpl *template.Template

func init() {
	navigationBarHTML = assets.MustAssetString("web/templates/navigation_bar.html")

	homepageHTML := assets.MustAssetString("web/templates/index.html")
	homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))

	secondViewHTML := assets.MustAssetString("web/templates/second_view.html")
	secondViewTpl = template.Must(template.New("second_view").Parse(secondViewHTML))

	// thirdViewFuncMap := ThirdViewFormattingFuncMap()
	// thirdViewHTML := assets.MustAssetString("templates/third_view.html")
	// thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))
}

type api struct {
	router     http.Handler
	repository task.TaskRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo task.TaskRepository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter() //creamos una instancia de router

	//endpoints:
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/tasks", a.fetchTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks", a.addTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{ID:[a-zA-Z0-9_]+}", a.removeTask).Methods(http.MethodDelete)
	r.HandleFunc("/tasks/{ID:[a-zA-Z0-9_]+}", a.modifyTask).Methods(http.MethodPut)
	r.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", http.FileServer(http.Dir("./web/static"))))
	//r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler { //metodo
	return a.router
}

// HomeHandler renders the homepage view template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/web/static/style.css")
	push(w, "/web/static/todolist.css")
	push(w, "/web/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

//Handlers:
func (a *api) fetchTasks(w http.ResponseWriter, r *http.Request) { //para mostrar todas las tareas
	tasks, _ := a.repository.FetchGophers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (a *api) addTask(w http.ResponseWriter, r *http.Request) { //para añadir nueva tarea
	decoder := json.NewDecoder(r.Body)

	var t task.Task
	decoder.Decode(&t)

	w.Header().Set("Content-Type", "application/json")

	a.repository.CreateGopher(&t)

	w.WriteHeader(http.StatusCreated)
}

func (a *api) removeTask(w http.ResponseWriter, r *http.Request) { //para borrar una tarea
	vars := mux.Vars(r)
	a.repository.DeleteGopher(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (a *api) modifyTask(w http.ResponseWriter, r *http.Request) { //para marcar una tarea como realizada
	decoder := json.NewDecoder(r.Body)

	var t task.Task
	decoder.Decode(&t)

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	a.repository.UpdateGopher(vars["ID"], &t)

	w.WriteHeader(http.StatusNoContent)
}
