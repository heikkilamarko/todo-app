FROM grafana/grafana-oss:9.3.11
RUN grafana-cli plugins install marcusolsson-json-datasource
COPY ./dashboards /etc/todo-app/dashboards
COPY ./provisioning /etc/todo-app/provisioning
