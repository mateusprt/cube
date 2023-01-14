package task

import (
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateObject(t *testing.T) {
	state := Pending
	port, _ := nat.NewPort("tcp", "8080")
	portMap := map[nat.Port]struct{}{port: struct{}{}}
	task := &Task{
		ID:            uuid.New(),
		Name:          "Test task",
		State:         state,
		Image:         "image",
		Memory:        1000,
		Disk:          10000,
		ExposedPorts:  portMap,
		PortBindings:  map[string]string{"80": "8080"},
		RestartPolicy: "Some policy",
		StartTime:     time.Now(),
		FinishTime:    time.Now(),
	}
	assert.NotEmpty(t, task.ID)
	assert.NotEmpty(t, task.Name)
	assert.Equal(t, task.State, state)
	assert.NotEmpty(t, task.Image)
	assert.NotEmpty(t, task.Memory)
	assert.NotEmpty(t, task.Disk)
	assert.Equal(t, len(task.ExposedPorts), 1)
	assert.Equal(t, len(task.PortBindings), 1)
	assert.NotEmpty(t, task.RestartPolicy)
	assert.NotEmpty(t, task.StartTime)
	assert.NotEmpty(t, task.FinishTime)
}

func TestTaskEvent(t *testing.T) {
	state := Pending
	task := &Task{
		ID:            uuid.New(),
		Name:          "Test task",
		State:         state,
		Image:         "image",
		Memory:        1000,
		Disk:          10000,
		ExposedPorts:  make(nat.PortSet),
		PortBindings:  map[string]string{"80": "8080"},
		RestartPolicy: "Some policy",
		StartTime:     time.Now(),
		FinishTime:    time.Now(),
	}

	tasEvent := &TaskEvent{
		ID:        uuid.New(),
		State:     task.State,
		Timestamp: time.Now(),
		Task:      *task,
	}

	assert.NotEmpty(t, tasEvent.ID)
	assert.Equal(t, tasEvent.State, state)
	assert.Equal(t, tasEvent.State, task.State)
	assert.NotEmpty(t, tasEvent.Timestamp)
	assert.NotNil(t, tasEvent.Task)
	assert.Equal(t, tasEvent.Task.ID, task.ID)
}
