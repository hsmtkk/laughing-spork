FROM golang:1.18 AS builder
WORKDIR /opt
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build

FROM gcr.io/distroless/cc-debian11 AS runtime
COPY --from=builder /opt/laughing-spork /usr/local/bin/laughing-spork
ENTRYPOINT ["/usr/local/bin/laughing-spork"]