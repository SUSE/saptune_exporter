package main

import (
	"testing"
)

func TestNoteCollector(t *testing.T) {
	collector, _ := NewNoteCollector("test/fake_saptune_note.sh")
	Metrics(t, collector, "note.metrics")
}
