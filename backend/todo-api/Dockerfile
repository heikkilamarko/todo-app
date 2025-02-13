FROM golang AS build
COPY . .
ENV GOPATH=""
ENV CGO_ENABLED=0
RUN go build -trimpath -a -ldflags="-w -s" ./cmd/todo-api

FROM gcr.io/distroless/static
COPY --from=build /go/todo-api /todo-api
ENTRYPOINT ["/todo-api"]
