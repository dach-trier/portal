FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine3.23 AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app

# --
# go.mod
# --

COPY go.mod go.sum ./
RUN  go mod download

COPY cmd/migration ./cmd/migration
RUN  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ./out/migration ./cmd/migration

# --
# runtime image
# --

FROM alpine:3.23
WORKDIR /app
COPY database/migrations /app/database/migrations
COPY --from=build /app/out/migration /app/migration
ENTRYPOINT ["/app/migration"]
