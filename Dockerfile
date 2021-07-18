FROM node:alpine-lts as client

WORKDIR /app

COPY client/package.json /app/package.json
COPY client/yarn.lock /app/yarn.lock

RUN npm ci

COPY client .

RUN npm build

FROM golang:1.16-stretch as server-base

WORKDIR /go/github.com/dasdachs/inventory

ENV CGO_ENABLED=0
ARG IS_ARM
ARG TARGETOS
ARG TARGETARCH
ARG ARMVERSION

COPY server .

FROM server-base as server-0

RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -o inventory

FROM server-base as server-1

RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} GOARM={ARMVERSION: -7} go build -o inventory

FROM scratch
COPY --from=client /app/build /srv/inventory/public
COPY --from=server-{IS_ARM}  /go/github.com/dasdachs/inventory/inventory /srv/inventory
ENTRYPOINT [ "/srv/inventory/inventory" ]
