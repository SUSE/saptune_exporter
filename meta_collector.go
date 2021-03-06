package main

import (
	"strconv"

	"github.com/SUSE/saptune/txtparser"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const subsystem_meta = "meta"

// MetaCollector is the saptune collector for general infos
type MetaCollector struct {
	DefaultCollector
	saptuneSysConf string
}

// NewMetaCollector creates a new solution saptune collector
func NewMetaCollector(saptuneSysConf string) (*MetaCollector, error) {
	c := &MetaCollector{
		NewDefaultCollector(subsystem_meta),
		saptuneSysConf,
	}
	// this metric are set by  setSolutionEnabledMetric
	c.SetDescriptor("version", "Show version of saptune", nil)

	return c, nil
}

// Collect various metrics for saptune solution
func (c *MetaCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting saptune solution metrics...")
	c.setSaptuneVersionMetric(ch)
}

func (c *MetaCollector) setSaptuneVersionMetric(ch chan<- prometheus.Metric) {
	// get major saptune version
	sconf, err := txtparser.ParseSysconfigFile(c.saptuneSysConf, true)
	if err != nil {
		log.Warnf("Error: Unable to read configuration %s: %v\n", c.saptuneSysConf, err)
	}
	SaptuneVersion := sconf.GetString("SAPTUNE_VERSION", "")
	if SaptuneVersionF, err := strconv.ParseFloat(SaptuneVersion, 32); err == nil {
		ch <- c.MakeGaugeMetric("version", float64(SaptuneVersionF))
	}

}
