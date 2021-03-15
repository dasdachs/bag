FROM golang:1.16-stretch as builder

WORKDIR /go/github.com/dasdachs/inventory

ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH

COPY server/ .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o inventory

FROM scratch
COPY --from=builder /go/github.com/dasdachs/inventory/inventory /usr/local/bin/inventory
ENTRYPOINT [ "/usr/local/bin/inventory" ]
