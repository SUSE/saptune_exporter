package main

import (
	"testing"

	"github.com/SUSE/saptune_exporter/internal"
)

func TestNoteCollector(t *testing.T) {
	collector, _ := NewNoteCollector("test/fake_saptune_note.sh")
	internal.Metrics(t, collector, "note.metrics")
}
