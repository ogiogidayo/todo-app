package database

import (
	"context"
	"github.com/ogiogidayo/todo-app/testutil/fixture"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ogiogidayo/todo-app/clock"
	"github.com/ogiogidayo/todo-app/domain"
	"github.com/ogiogidayo/todo-app/testutil"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	t.Cleanup(func() {
		tx.Rollback()
	})
	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Fatalf("differs: (-got +want)\n%s", d)
	}
}

func prepareUser(ctx context.Context, t *testing.T, db Execer) domain.UserID {
	t.Helper()
	u := fixture.User(nil)
	result, err := db.ExecContext(ctx,
		`INSERT INTO user (name, password, role, created, modified)
				VALUES (?, ?, ?, ?, ?);`,
		u.Name, u.Password, u.Role, u.Created, u.Modified)
	if err != nil {
		t.Fatalf("insert user: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("got user_id: %v", err)
	}
	return domain.UserID(id)
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) (domain.UserID, domain.Tasks) {
	t.Helper()
	userID := prepareUser(ctx, t, con)
	otherUserID := prepareUser(ctx, t, con)
	c := clock.FixedClocker{}
	wants := domain.Tasks{
		{
			UserID: userID,
			Title:  "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			UserID: userID,
			Title:  "want task 2", Status: "done",
			Created: c.Now(), Modified: c.Now(),
		},
	}
	tasks := domain.Tasks{
		wants[0],
		{
			UserID: otherUserID,
			Title:  "not want task", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		wants[1],
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO task (user_id, title, status, created, modified)
			VALUES
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?);`,
		tasks[0].UserID, tasks[0].Title, tasks[0].Status, tasks[0].Created, tasks[0].Modified,
		tasks[1].UserID, tasks[1].Title, tasks[1].Status, tasks[1].Created, tasks[1].Modified,
		tasks[2].UserID, tasks[2].Title, tasks[2].Status, tasks[2].Created, tasks[2].Modified,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	tasks[0].ID = domain.TaskID(id)
	tasks[1].ID = domain.TaskID(id + 1)
	tasks[2].ID = domain.TaskID(id + 2)
	return userID, wants
}
