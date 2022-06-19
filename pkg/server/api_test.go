package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jorgefg4/todolist/pkg/data"
	"github.com/jorgefg4/todolist/pkg/database"
	"github.com/jorgefg4/todolist/pkg/task"
)

func TestFetchTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	repo := database.NewTaskRepository(data.Tasks)
	s := New(repo)

	rec := httptest.NewRecorder() //con el paquete httptest podemos generar el http.ResponseWriter

	s.fetchTasks(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	var got []task.Task
	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	expected := len(data.Tasks)

	if len(got) != expected {
		t.Errorf("expected %v tasks, got: %v task", data.Tasks, got)
	}
}
