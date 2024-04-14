package handler

import (
	"net/http"

	"github.com/ogiogidayo/todo-app/domain"
)

type ListTask struct {
	Sevices ListTaskService
}

type task struct {
	ID     domain.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status domain.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Sevices.ListTasks(ctx)
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
