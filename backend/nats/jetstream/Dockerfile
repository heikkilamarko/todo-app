FROM synadia/nats-box

COPY . /

RUN chmod +x /configure_jetstream.sh

ENTRYPOINT [ "/configure_jetstream.sh" ]
