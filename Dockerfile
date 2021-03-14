FROM golang:1.16-stretch as builder

WORKDIR /go/github.com/dasdachs/inventory

COPY server/ .

RUN go build -o inventory

# FROM scratch
# COPY --from=builder /go/github.com/dasdachs/inventory/inventory /usr/local/bin/inventory
# COPY --from=builder /etc/passwd /etc/passwd
# ENTRYPOINT [ "/usr/local/bin/inventory" ]
