package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	EventsNotFoundTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_car_api_events_not_found_total",
		Help: "Total number of not found events",
	})

	EventsCUDTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_car_api_events_create_update_delete_total",
		Help: "Total number of CUD events",
	})

	EventsProcessedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_car_api_events_processed_total",
		Help: "Total number of processed events",
	})
)
