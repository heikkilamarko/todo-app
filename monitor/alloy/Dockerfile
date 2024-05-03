FROM grafana/alloy
COPY ./config.alloy /etc/alloy/config.alloy
CMD ["run", "--storage.path=/var/lib/alloy/data", "/etc/alloy/config.alloy"]
