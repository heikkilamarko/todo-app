FROM alpine

RUN apk --no-cache add curl
RUN apk --no-cache add gettext

COPY . /

RUN chmod +x /create_grafana_resources.sh

ENTRYPOINT [ "/create_grafana_resources.sh" ]
