1. Run go-metric-producer +prometheus +grafana by
$ docker-compose up
2. Visit grafana on http://localhost:3000/
3. 'Configuration' -> 'Data Sources' -> 'Prometheus' -> Set 'URL' field to 'http://prometheus:9090' -> Press 'Save and test'
4. 'Create' -> 'Dashboard' -> 'Add new panel' -> 'Metrics browser' -> Select 'promtestapp_current_wait_period' or 'promtestapp_operations_processed_total'
