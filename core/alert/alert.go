package alert

import (
	"context"
	"time"
)

type Repository interface {
	Create(context.Context, Alert) (Alert, error)
	List(context.Context, Filter) ([]Alert, error)
	BulkUpdateSilence(context.Context, []int64, string) error
}

type Alert struct {
	ID            uint64    `json:"id"`
	ProviderID    uint64    `json:"provider_id"`
	NamespaceID   uint64    `json:"namespace_id"`
	ResourceName  string    `json:"resource_name"`
	MetricName    string    `json:"metric_name"`
	MetricValue   string    `json:"metric_value"`
	Severity      string    `json:"severity"`
	Rule          string    `json:"rule"`
	TriggeredAt   time.Time `json:"triggered_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	SilenceStatus string    `json:"silence_status"`

	// These fields won't be stored in the DB
	// these are additional information for notification purposes
	GroupKey     string
	Status       string
	Annotations  map[string]string
	Labels       map[string]string
	GeneratorURL string
	Fingerprint  string
}
