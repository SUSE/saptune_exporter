package main

import (
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const subsystemNote = "note"

// NoteCollector is the saptune solution collector
type NoteCollector struct {
	DefaultCollector
}

// NewNoteCollector creates a new solution saptune collector
func NewNoteCollector() (*NoteCollector, error) {
	c := &NoteCollector{
		NewDefaultCollector(subsystemNote),
	}

	c.SetDescriptor("enabled", "This metrics show with 1 all the enabled notes on the system", []string{"noteID"})

	return c, nil
}

// Collect various metrics for saptune solution
func (c *NoteCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting saptune note metrics...")
	c.setNoteListMetric(ch)
}

func (c *NoteCollector) setNoteListMetric(ch chan<- prometheus.Metric) {
	noteList, err := exec.Command("saptune", "note", "enabled").CombinedOutput()
	if err != nil {
		log.Warnf("%v - Failed to run saptune note enabled command n %s ", err, string(noteList))
		return
	}

	notes := strings.Fields(string(noteList))

	for _, note := range notes {
		ch <- c.MakeGaugeMetric("enabled", float64(1), note)
	}
}
