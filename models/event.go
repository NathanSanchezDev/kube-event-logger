package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Type        int64     `json:"type"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    string    `json:"severity"`
	Source      string    `json:"source"`
	ClusterName string    `json:"clusterName"`
	Namespace   string    `json:"namespace"`
}
