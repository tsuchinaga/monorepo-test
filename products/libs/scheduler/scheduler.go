package scheduler

import (
	"github.com/robfig/cron/v3"
)

type Scheduler interface {
	Start() error
	Stop()
	AddJob(name string, spec string, job Job)
}

// NewScheduler - スケジューラの生成
func NewScheduler() Scheduler {
	return &scheduler{cron: cron.New()}
}

type scheduler struct {
	cron *cron.Cron
	jobs []struct {
		name    string
		spec    string
		job     Job
		entryId cron.EntryID
	}
}

func (s *scheduler) Start() error {
	for _, job := range s.jobs {
		if entryId, err := s.cron.AddJob(job.spec, job.job); err != nil {
			return err
		} else {
			job.entryId = entryId
		}
	}
	s.cron.Start()
	return nil
}

func (s *scheduler) Stop() {
	s.cron.Stop().Done()
}

func (s *scheduler) AddJob(name string, spec string, job Job) {
	s.jobs = append(s.jobs, struct {
		name    string
		spec    string
		job     Job
		entryId cron.EntryID
	}{name: name, spec: spec, job: job})
}

type Job interface {
	Run()
}
