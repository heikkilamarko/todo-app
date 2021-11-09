FROM grafana/promtail:2.4.1

COPY ./promtail.yml /etc/promtail/config.yml

CMD ["-config.file=/etc/promtail/config.yml", "-config.expand-env=true"]
