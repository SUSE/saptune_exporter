package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const saptunePath = "/usr/sbin/saptune"

func main() {

	// register various collectors
	solutionCollector, err := NewSolutionCollector()
	if err != nil {
		log.Warn(err)
	} else {
		prometheus.MustRegister(solutionCollector)
		log.Info("Saptune Solution collector registered")
	}
	metaCollector, err := NewMetaCollector()
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
	http.HandleFunc("/", landing)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Serving metrics on port 9758")
	log.Fatal(http.ListenAndServe(":9758", nil))
}

func landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<html>
<head>
	<title>SUSE  Saptune Exporter</title>
</head>
<body>
	<h1>SUSE Saptune expoter</h1>
	<h2>Prometheus exporter for Saptune</h2>
	<ul>
		<li><a href="metrics">Metrics</a></li>
		<li><a href="https://github.com/SUSE/saptune" target="_blank">GitHub</a></li>
	</ul>
</body>
</html>
`))
}
