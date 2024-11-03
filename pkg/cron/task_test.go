package cron

import (
	"testing"
	"time"
)

type DummyTask struct {
}

func (d *DummyTask) Run() {
	println("dummy")

}

func TestTask(t *testing.T) {
	InitCron([]*Task{NewTask("dummy", "*/1 * * * * *", new(DummyTask))})
	time.Sleep(20 * time.Hour)
}
