package model

import (
	"time"

	"github.com/odpf/siren/core/alert"
)

type Alert struct {
	ID           uint64    `db:"id"`
	ProviderID   uint64    `db:"provider_id"`
	ResourceName string    `db:"resource_name"`
	MetricName   string    `db:"metric_name"`
	MetricValue  string    `db:"metric_value"`
	Severity     string    `db:"severity"`
	Rule         string    `db:"rule"`
	TriggeredAt  time.Time `db:"triggered_at"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (a *Alert) FromDomain(alrt alert.Alert) {
	a.ID = alrt.ID
	a.ProviderID = alrt.ProviderID
	a.ResourceName = alrt.ResourceName
	a.MetricName = alrt.MetricName
	a.MetricValue = alrt.MetricValue
	a.Severity = alrt.Severity
	a.Rule = alrt.Rule
	a.TriggeredAt = alrt.TriggeredAt
	a.CreatedAt = alrt.CreatedAt
	a.UpdatedAt = alrt.UpdatedAt
}

func (a *Alert) ToDomain() *alert.Alert {
	return &alert.Alert{
		ID:           a.ID,
		ProviderID:   a.ProviderID,
		ResourceName: a.ResourceName,
		MetricName:   a.MetricName,
		MetricValue:  a.MetricValue,
		Severity:     a.Severity,
		Rule:         a.Rule,
		TriggeredAt:  a.TriggeredAt,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}