package queue

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"time"
)

type CronJob struct {
	Timestamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	Enabled   bool      `json:"enabled"`
	TTL       int       `json:"interval"`
}

func NewCronJob(action string, enabled bool, ttl int) *CronJob {
	return &CronJob{
		Timestamp: time.Now(),
		Action:    action,
		Enabled:   enabled,
		TTL:       ttl,
	}
}

func (c *CronJob) Encode() []byte {
	payload, err := json.Marshal(c)
	if err != nil {
		log.Err(err).Msg("failed to marshal cronjob")
		return nil
	}
	return payload
}
