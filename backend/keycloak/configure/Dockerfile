FROM alpine
RUN apk --no-cache add curl
RUN apk --no-cache add jq
COPY . /
RUN chmod +x /configure.sh
ENTRYPOINT [ "/configure.sh" ]
