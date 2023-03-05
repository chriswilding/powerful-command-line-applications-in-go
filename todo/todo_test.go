package todo_test

import (
	"os"
	"testing"

	"github.com/chriswilding/powerful-command-line-applications-in-go/todo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	assert.Equal(t, taskName, l[0].Task)
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	assert.Equal(t, taskName, l[0].Task)
	assert.False(t, l[0].Done)

	l.Complete(1)

	assert.True(t, l[0].Done)
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	assert.Equal(t, tasks[0], l[0].Task)

	l.Delete(2)

	assert.Len(t, l, 2)
	assert.Equal(t, l[1].Task, tasks[2])
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	assert.Equal(t, l1[0].Task, taskName)

	tf, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(tf.Name())

	err = l1.Save(tf.Name())
	require.NoError(t, err)

	err = l2.Get(tf.Name())
	require.NoError(t, err)

	assert.Equal(t, l1[0].Task, l2[0].Task)
}
