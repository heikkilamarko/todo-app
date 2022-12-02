FROM grafana/promtail:2.7.0
COPY ./promtail.yml /etc/promtail/config.yml
CMD ["-config.file=/etc/promtail/config.yml", "-config.expand-env=true"]
