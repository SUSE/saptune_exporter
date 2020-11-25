package internal

import "net/http"

// Landing just write the HELP message for thee main / handler
func Landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<html>
<head>
	<title>SUSE  Saptune Exporter</title>
</head>
<body>
	<h1>SUSE Saptune exporter</h1>
	<h2>Prometheus exporter for Saptune</h2>
	<ul>
		<li><a href="metrics">Metrics</a></li>
		<li><a href="https://github.com/SUSE/saptune_exporter" target="_blank">GitHub</a></li>
	</ul>
</body>
</html>
`))
}
