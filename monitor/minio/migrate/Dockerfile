FROM minio/mc

COPY . /

RUN chmod +x /migrate.sh

ENTRYPOINT [ "/migrate.sh" ]
