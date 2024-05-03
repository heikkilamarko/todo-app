FROM grafana/promtail:3.0.0
COPY ./promtail.yml /etc/promtail/config.yml
CMD ["-config.file=/etc/promtail/config.yml", "-config.expand-env=true"]
