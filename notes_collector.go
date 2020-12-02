package main

import (
	"os/exec"
	"strings"

	"github.com/SUSE/saptune/sap/note"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const subsystemNote = "note"
const noteTuningSheets = "/usr/share/saptune/notes/"
const extraTuningSheets = "/etc/saptune/extra/"

// NoteCollector is the saptune solution collector
type NoteCollector struct {
	DefaultCollector
	saptunePath string
}

// Lookup for a given Solution ID enabled the corrisponding Name/Descriptions
func getNoteDesc(enabledNoteID string) string {
	tuningOptions := note.GetTuningOptions(noteTuningSheets, extraTuningSheets)
	for _, noteID := range tuningOptions.GetSortedIDs() {
		noteObj := tuningOptions[noteID]
		if noteID == enabledNoteID {
			// this is how looks like noteObj.Name, remove the version.
			// "Linux: User and system resource limits \n			Version 5 from 18.06.2018 ",
			solDescRaw := strings.Split(noteObj.Name(), "\n")
			// return only description
			return solDescRaw[0]
		}
	}
	// in case there is no match return empty string
	return ""
}

// NewNoteCollector creates a new solution saptune collector
func NewNoteCollector(saptunePath string) (*NoteCollector, error) {
	c := &NoteCollector{
		NewDefaultCollector(subsystemNote),
		saptunePath,
	}

	c.SetDescriptor("enabled", "This metrics show with 1 all the enabled notes on the system", []string{"note_id", "note_desc"})

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
		noteDesc := getNoteDesc(note)
		if noteDesc == "" {
			log.Warnf("Could not find the note description for given note ID %s", note)
		}
		ch <- c.MakeGaugeMetric("enabled", float64(1), note, noteDesc)
	}
}
