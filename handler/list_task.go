package handler

import (
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/entity"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo *database.Repository
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
