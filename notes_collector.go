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
	saptunePath string
}

// NewNoteCollector creates a new solution saptune collector
func NewNoteCollector(saptunePath string) (*NoteCollector, error) {
	c := &NoteCollector{
		NewDefaultCollector(subsystemNote),
		saptunePath,
	}

	c.SetDescriptor("enabled", "This metrics show with 1 all the enabled notes on the system", []string{"note_id"})

	return c, nil
}

// Collect various metrics for saptune solution
func (c *NoteCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting saptune note metrics...")
	c.noteEnabled(ch)
}

func (c *NoteCollector) noteEnabled(ch chan<- prometheus.Metric) {
	err := checkExecutables(c.saptunePath)
	if err != nil {
		log.Warnf("%v failed to retrieve saptune executable", err)
		return
	}
	noteList, err := exec.Command(c.saptunePath, "note", "enabled").CombinedOutput()
	if err != nil {
		log.Warnf("%v - Failed to run saptune note enabled command n %s ", err, string(noteList))
		return
	}

	notes := strings.Fields(string(noteList))

	for _, note := range notes {
		ch <- c.MakeGaugeMetric("enabled", float64(1), note)
	}
}
