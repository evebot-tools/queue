package queue

import (
	"time"
)

type CronJob struct {
	Timestamp  time.Time `json:"timestamp"`
	Action     string    `json:"action"`
	Enabled    bool      `json:"enabled"`
	TTL        int       `json:"interval"`
	StringData *string   `json:"stringData,omitempty"`
	IntData    *int      `json:"intData,omitempty"`
}

func NewCronJob(action string, enabled bool, ttl int) *CronJob {
	return &CronJob{
		Timestamp:  time.Now(),
		Action:     action,
		Enabled:    enabled,
		TTL:        ttl,
		StringData: nil,
		IntData:    nil,
	}
}

func NewCronJobWithStringData(action string, enabled bool, ttl int, data string) *CronJob {
	return &CronJob{
		Timestamp:  time.Now(),
		Action:     action,
		Enabled:    enabled,
		TTL:        ttl,
		StringData: &data,
		IntData:    nil,
	}
}

func NewCronJobWithIntData(action string, enabled bool, ttl int, data int) *CronJob {
	return &CronJob{
		Timestamp:  time.Now(),
		Action:     action,
		Enabled:    enabled,
		TTL:        ttl,
		StringData: nil,
		IntData:    &data,
	}
}
