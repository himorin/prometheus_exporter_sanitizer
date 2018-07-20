### prometheus_exporter_sanitizer

This is a sanitizer proxy module for prometheus exporter. 
Some exporter (like JIRA exporter) sometimes output two sections for one 
metrics, which cause parsing error at prometheus on second line for metrics 
names by having two lines of HELP or TYPE.

This module works as proxy between prometheus and original exporter, 
gets specified metrics exporter and returns sanitized lines when this receives 
request(s). 


### Parameters

This module accepts following command line parameters.

- port: port to expose /metrics on, default ":9432"
- origin: URL of original exporter, default "http://localhost:9100/metrics"

