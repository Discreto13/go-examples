FROM golang:1.18.1 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd .
RUN go build -o metric-producer

ENTRYPOINT ["/app/metric-producer"]