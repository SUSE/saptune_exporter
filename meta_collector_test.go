package main

import (
	"testing"

	"github.com/SUSE/saptune_exporter/internal"
)

func TestMetaCollector(t *testing.T) {
	collector, _ := NewMetaCollector("test/fake_saptune_meta.sh")
	internal.Metrics(t, collector, "meta.metrics")
}
