FROM node:lts-alpine AS client

WORKDIR /app

ENV NODE_ENV production

COPY ./package.json ./package.json
COPY ./pnpm.lock ./yarn.lock

RUN pnpm install --frozen-lockfile --production

COPY . .

RUN pnpm lint && pnpm test && pnpm build


FROM golang:1.16.5-alpine3.13 AS server

WORKDIR /go/github.com/dasdachs/inventory

ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH

COPY --from=builder /app/build /src/_public

COPY server/go.mod .
COPY server/ .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o inventory

FROM scratch

WORKDIR /srv/inventory

COPY --from=server /go/github.com/dasdachs/inventory/inventory /srv/inventory

ENTRYPOINT [ "/srv/inventory" ]
