package domain

import "time"

type Alerts struct {
	Alerts []Alert `json:"alerts"`
}

type Labels struct {
	Severity string `json:"severity"`
}

type Annotations struct {
	Resource    string `json:"resource"`
	Template    string `json:"template"`
	MetricName  string `json:"metricName"`
	MetricValue string `json:"metricValue"`
}

type Alert struct {
	Status      string      `json:"status"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
}

type AlertHistoryObject struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	TemplateID  string    `json:"template_id"`
	MetricName  string    `json:"metric_name"`
	MetricValue string    `json:"metric_value"`
	Level       string    `json:"level"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type AlertHistoryService interface {
	Create(*Alerts) ([]AlertHistoryObject, error)
	Get(string, uint32, uint32) ([]AlertHistoryObject, error)
	Migrate() error
}