package main

import (
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/mateusprt/cube/manager"
	"github.com/mateusprt/cube/node"
	"github.com/mateusprt/cube/task"
	"github.com/mateusprt/cube/worker.go"
)

func main() {

	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]task.Task),
	}
	fmt.Printf("worker: %v\n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string]task.Task),
		EventDb: make(map[string]task.TaskEvent),
		//Workers: []string{w.Name},
		Workers: []string{"worker-1"},
	}
	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name: "Node-1",
		Ip:   "192.168.1.1",
		//Cores:  4,
		Memory: 1024,
		Disk:   25,
		//Role:   "worker",
	}
	fmt.Printf("node: %v\n", n)
}
