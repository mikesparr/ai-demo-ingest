FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/mikesparr/ai-demo-ingest/
WORKDIR /go/src/github.com/mikesparr/ai-demo-ingest
RUN go mod download
COPY . /go/src/github.com/mikesparr/ai-demo-ingest
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/ai-demo-ingest github.com/mikesparr/ai-demo-ingest

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/mikesparr/ai-demo-ingest/build/ai-demo-ingest /usr/bin/ai-demo-ingest
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/ai-demo-ingest"]