package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/cube2222/Blog/NATS/MasterWorker"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
	"github.com/satori/go.uuid"
)

var Tasks []Transport.Task
var TaskMutex sync.Mutex
var oldestFinishedTaskPointer int
var nc *nats.Conn

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments. Need NATS server address.")
		return
	}

	var err error

	nc, err = nats.Connect(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	Tasks = make([]Transport.Task, 0, 20)
	TaskMutex = sync.Mutex{}
	oldestFinishedTaskPointer = 0

	initTestTasks()

	nc.Subscribe("Work.TaskToDo", func(m *nats.Msg) {
		myTaskPointer, ok := getNextTask()
		if ok {
			data, err := proto.Marshal(myTaskPointer)
			if err == nil {
				nc.Publish(m.Reply, data)
			}
		}
	})

	nc.Subscribe("Work.TaskFinished", func(m *nats.Msg) {
		myTask := Transport.Task{}
		err := proto.Unmarshal(m.Data, &myTask)
		if err == nil {
			TaskMutex.Lock()
			Tasks[myTask.Id].State = 2
			Tasks[myTask.Id].Finisheduuid = myTask.Finisheduuid
			TaskMutex.Unlock()
		}
	})

	select {}
}

func getNextTask() (*Transport.Task, bool) {
	TaskMutex.Lock()
	defer TaskMutex.Unlock()
	for i := oldestFinishedTaskPointer; i < len(Tasks); i++ {
		if i == oldestFinishedTaskPointer && Tasks[i].State == 2 {
			oldestFinishedTaskPointer++
		} else {
			if Tasks[i].State == 0 {
				Tasks[i].State = 1
				go resetTaskIfNotFinished(i)
				return &Tasks[i], true
			}
		}
	}
	return nil, false
}

func resetTaskIfNotFinished(i int) {
	time.Sleep(2 * time.Minute)
	TaskMutex.Lock()
	if Tasks[i].State != 2 {
		Tasks[i].State = 0
	}
}

func initTestTasks() {
	for i := 0; i < 20; i++ {
		newTask := Transport.Task{Uuid: uuid.NewV4().String(), State: 0}
		fileServerAddressTransport := Transport.DiscoverableServiceTransport{}
		msg, err := nc.Request("Discovery.FileServer", nil, 1000*time.Millisecond)
		if err == nil && msg != nil {
			err := proto.Unmarshal(msg.Data, &fileServerAddressTransport)
			if err != nil {
				continue
			}
		}
		if err != nil {
			continue
		}

		fileServerAddress := fileServerAddressTransport.Address
		data := make([]byte, 0, 1024)
		buf := bytes.NewBuffer(data)
		fmt.Fprint(buf, "get,my,data,my,get,get,have")
		r, err := http.Post(fileServerAddress+"/"+newTask.Uuid, "", buf)
		if err != nil || r.StatusCode != http.StatusOK {
			continue
		}

		newTask.Id = int32(len(Tasks))
		Tasks = append(Tasks, newTask)
	}
}
