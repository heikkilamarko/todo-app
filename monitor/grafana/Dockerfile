FROM grafana/grafana-oss
RUN grafana-cli plugins install marcusolsson-json-datasource
COPY ./dashboards /etc/todo-app/dashboards
COPY ./provisioning /etc/todo-app/provisioning
