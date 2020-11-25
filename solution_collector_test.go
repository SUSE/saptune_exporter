package main

import (
	"testing"

	"github.com/SUSE/saptune_exporter/internal"
)

func TestSolutionEnabledCollector(t *testing.T) {
	collector, _ := NewSolutionCollector("test/fake_saptune_solution.sh")
	internal.Metrics(t, collector, "solution.metrics")
}
