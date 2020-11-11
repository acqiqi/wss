package task

import (
	"github.com/astaxie/beego/toolbox"
	"log"
)

type Task struct {
	TaskName string
	Spec     string
	Func     toolbox.TaskFunc
}

var Tk []Task

func init() {
	var tk = make([]Task, 1)
	tk[0] = Task{TaskName: "message_queue", Spec: "0/1 * * * * *", Func: MessagePushQueue}
	Tk = tk
	run()
}

func run() {
	log.Println("Run Task")
	for _, v := range Tk {
		tkobj := toolbox.NewTask(v.TaskName, v.Spec, v.Func)
		toolbox.AddTask(v.TaskName, tkobj)
	}
	toolbox.StartTask()
}
