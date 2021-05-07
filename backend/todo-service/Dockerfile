FROM golang:1.16.4-alpine3.13 AS build

COPY . .

ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" -o todo-service

FROM scratch

COPY --from=build /go/todo-service /todo-service

ENTRYPOINT ["/todo-service"]