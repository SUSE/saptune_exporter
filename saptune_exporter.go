package main

import (
	"net/http"

	"github.com/SUSE/saptune_exporter/internal"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const (
	saptunePath    = "/usr/sbin/saptune"
	saptuneSycConf = "/etc/sysconfig/saptune"
)

func main() {

	// register various collectors
	solutionCollector, err := NewSolutionCollector(saptunePath)
	if err != nil {
		log.Warn(err)
	} else {
		prometheus.MustRegister(solutionCollector)
		log.Info("Saptune Solution collector registered")
	}
	metaCollector, err := NewMetaCollector(saptuneSycConf)
	if err != nil {
		log.Warn(err)
	} else {
		prometheus.MustRegister(metaCollector)
		log.Info("Saptune Meta collector registered")
	}

	noteCollector, err := NewNoteCollector(saptunePath)
	if err != nil {
		log.Warn(err)
	} else {
		prometheus.MustRegister(noteCollector)
		log.Info("Saptune Note collector registered")
	}

	// disable golang specific metrics
	prometheus.Unregister(prometheus.NewGoCollector())

	// serve metrics
	http.HandleFunc("/", internal.Landing)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Serving metrics on port 9758")
	log.Fatal(http.ListenAndServe(":9758", nil))
}
