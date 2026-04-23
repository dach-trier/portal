FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine3.23 AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app

# --
# packages & pnpm
# --

RUN apk add --no-cache npm
RUN npm install -g pnpm

# --
# go.mod
# --

COPY go.mod go.sum ./
RUN  go mod download

# --
# source code & styles & assets & db
# --

COPY package.json ./package.json
COPY pnpm-lock.yaml ./pnpm-lock.yaml
RUN  pnpm install

COPY database ./database
COPY sqlc.yml ./sqlc.yml
RUN  go tool sqlc generate

COPY templates ./templates
COPY icons ./icons
COPY i18n ./i18n

COPY tailwind.css ./tailwind.css
COPY styles ./styles
COPY third_party/styles/button.tailwind.css ./third_party/styles/button.tailwind.css
RUN  pnpm exec tailwindcss -i ./tailwind.css -o ./assets/styles.css --minify

COPY internal ./internal
COPY cmd/server ./cmd/server
RUN  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ./out/server ./cmd/server

# --
# runtime image
# --

FROM alpine:3.23
WORKDIR /app
COPY --from=build /app/out/server /app/server
COPY --from=build /app/assets     /app/assets
COPY --from=build /app/templates  /app/templates
COPY --from=build /app/icons      /app/icons
COPY --from=build /app/i18n       /app/i18n
ENTRYPOINT ["/app/server"]
