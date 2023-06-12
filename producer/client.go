package main

import (
	"asynq-demo/task"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	task, err := NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	/**
		1. 指定消费时间 asynq.ProcessAt(time.Now())
		2.指定过多久消费 asynq.ProcessIn(2*time.Second)
	    3.  指定队列名 asynq.Queue("")
	    4. 指定重试次数 asynq.MaxRetry()
	    5. 指定消费后保存多久 asynq.Retention(30*time.Minute)
	*/
	info, err := client.Enqueue(task, asynq.ProcessAt(time.Now()))

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

}

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(task.EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(task.TypeEmailDelivery, payload), nil
}
