FROM golang:1.14.7-alpine

WORKDIR /resque_exporter

COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/resque_exporter/resque_exporter.go

EXPOSE 5555

CMD ["./resque_exporter","--config","config.yaml"]