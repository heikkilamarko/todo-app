FROM grafana/loki:2.8.2
COPY ./loki.yml /etc/loki/config.yml
CMD ["-config.file=/etc/loki/config.yml", "-config.expand-env=true"]
