package main

import (
	"log"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
)

func main() {
	sd, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: "carolyntestprj-262400",
		// ReportingInterval sets the frequency of reporting metrics
		// to stackdriver backend.
		ReportingInterval: 60 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create the Stackdriver exporter: %v", err)
	}
	// It is imperative to invoke flush before your main function exits
	defer sd.Flush()

	// Start the metrics exporter
	sd.StartMetricsExporter()
	defer sd.StopMetricsExporter()
}
