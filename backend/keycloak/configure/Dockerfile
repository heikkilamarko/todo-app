FROM alpine
RUN apk --no-cache add curl jq
COPY . /
RUN chmod +x /configure.sh
ENTRYPOINT [ "/configure.sh" ]
