package cron

import (
	cron "github.com/robfig/cron/v3"
)

var c *cron.Cron

var tasks map[string]Task = make(map[string]Task)

type Task struct {
	Name        string
	CronExpress string
	Job         cron.Job
	EntryId     cron.EntryID
}

func NewTask(name string, cronExpress string, job cron.Job) *Task {
	return &Task{
		Name:        name,
		CronExpress: cronExpress,
		Job:         job,
	}
}

func InitCron(ts []*Task) error {
	c = cron.New(cron.WithSeconds(), cron.WithLogger(cron.DefaultLogger), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	go c.Run()
	for i := range ts {
		var err error
		ts[i].EntryId, err = c.AddJob(ts[i].CronExpress, ts[i].Job)
		if err != nil {
			return err
		}
		tasks[ts[i].Name] = *ts[i]
	}
	return nil
}

func AddTask(ts []*Task) error {
	for i := range ts {
		var err error
		ts[i].EntryId, err = c.AddJob(ts[i].CronExpress, ts[i].Job)
		tasks[ts[i].Name] = *ts[i]
		if err != nil {
			return err
		}
	}
	return nil
}
