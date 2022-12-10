FROM grafana/loki:2.7.1
COPY ./loki.yml /etc/loki/config.yml
CMD ["-config.file=/etc/loki/config.yml", "-config.expand-env=true"]
