FROM grafana/loki:2.2.1

COPY ./loki.yml /etc/loki/config.yml

CMD ["-config.file=/etc/loki/config.yml", "-config.expand-env=true"]
