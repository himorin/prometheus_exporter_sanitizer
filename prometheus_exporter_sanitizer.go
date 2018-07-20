/* Reformatter for prometheus exporter
 */

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	port      = flag.String("port", ":9432", "port to expose /metrics on")
	origin    = flag.String("origin", "http://localhost:9100/metrics", "URL of original exporter")
)

func FilterLines (data []byte) string {
	var (
		ret []string
		lines map[string]int16
	)
	lines = map[string]int16{}
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			if _, ok := lines[line]; ok {
				continue
			}
			lines[line] = 0
		}
		ret = append(ret, line)
	}
	return strings.Join(ret, "\n") + "\n"
}

func main () {
	flag.Parse()
	http.HandleFunc("/metrics", func (w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(*origin)
		if err != nil {
			w.Write([]byte(`# error`))
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				w.Write([]byte(`# error`))
			} else {
				w.Write([]byte(FilterLines(body)))
			}
		}
	})
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head><title>exporter formatter</title></head>
			<body><a href="/metrics">Metrics</a></body></html>`))
	})
	log.Fatal(http.ListenAndServe(*port, nil))
}

