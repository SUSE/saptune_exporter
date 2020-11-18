package main

import (
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const subsystem = "solution"

// SolutionCollector is the saptune solution collector
type SolutionCollector struct {
	DefaultCollector
}

// NewSolutionCollector creates a new solution saptune collector
func NewSolutionCollector() (*SolutionCollector, error) {
	c := &SolutionCollector{
		NewDefaultCollector(subsystem),
	}
	c.SetDescriptor("enabled", "show the enabled solution's name. 1 means is enabled. disabled metric is absent ", []string{"solutionName"})
	c.SetDescriptor("compliant", "show if current solution applied is compliant 1 or not 0", []string{"solutionName"})
	return c, nil
}

// Collect various metrics for saptune solution
func (c *SolutionCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting saptune solution metrics...")

	// solution enabled
	solutionName, err := exec.Command("saptune", "solution", "enabled").CombinedOutput()
	if err != nil {
		log.Warnf("%v - Failed to run saptune solution enabled command n %s ", err, string(solutionName))
		return
	}
	ch <- c.MakeGaugeMetric("enabled", float64(1), string(solutionName))

	// TODO: the return code is a "fragile" check to base the metrics up on this
	// is something could be improve on the saptune CLI
	_, err = exec.Command("saptune", "solution", "verify").CombinedOutput()
	if err != nil {
		ch <- c.MakeGaugeMetric("compliant", float64(0), string(solutionName))
		return
	}
	// no error so the solution is compliant
	ch <- c.MakeGaugeMetric("compliant", float64(1), string(solutionName))
}
