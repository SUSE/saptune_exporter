package main

import (
	"testing"
)

func TestMetaCollector(t *testing.T) {
	collector, _ := NewMetaCollector("test/fake_saptune_meta.sh")
	Metrics(t, collector, "meta.metrics")
}
