FROM quay.io/keycloak/keycloak AS builder
ENV KC_DB=postgres
RUN /opt/keycloak/bin/kc.sh build

FROM quay.io/keycloak/keycloak
COPY --from=builder /opt/keycloak/lib/quarkus/ /opt/keycloak/lib/quarkus/
ENTRYPOINT ["/opt/keycloak/bin/kc.sh", "start"]
