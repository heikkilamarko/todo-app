FROM golang:1.24-rc AS build
COPY . .
ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" ./cmd/todo-worker

FROM gcr.io/distroless/static
COPY --from=build /go/todo-worker /todo-worker
ENTRYPOINT ["/todo-worker"]
