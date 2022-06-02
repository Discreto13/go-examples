# What is it?

This is an example how metric services can be implemented in go project.

Used tools: prometheus+grafana

# How to run?

1. Run **'docker-compose up'**. It starts follow services:
   - go-metric-producer
   - prometheus
   - grafana
2. Open Grafana in browser (http://localhost:3000/).
   1. Choose **'Configuration'** on the left panel -> **'Data Sources'** -> **'Prometheus'** -> Set **'URL'** field with *'http://prometheus:9090/'* (specified in *docker-compose.yml*) -> Press **'Save and test'**
   2. Choose **'Create'**  on the left panel -> **'Dashboard'** -> **'Add new panel'** -> **'Metrics browser'** -> Select *'promtestapp_current_wait_period'* or *'promtestapp_operations_processed_total'* (custom metrics added in *main.go*)
