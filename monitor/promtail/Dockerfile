FROM grafana/promtail:2.8.2
COPY ./promtail.yml /etc/promtail/config.yml
CMD ["-config.file=/etc/promtail/config.yml", "-config.expand-env=true"]
