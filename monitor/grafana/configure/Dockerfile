FROM alpine
RUN apk --no-cache add curl
RUN apk --no-cache add gettext
COPY . /
RUN chmod +x /configure.sh
ENTRYPOINT [ "/configure.sh" ]
