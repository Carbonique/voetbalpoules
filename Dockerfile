FROM golang:1.18-alpine AS build
ARG version

WORKDIR /build
COPY ./ .

RUN go mod download

RUN CGO_ENABLED=0 go build --ldflags="-X 'github.com/Carbonique/voetbalpoules/cmd.Version=${version}'"

FROM alpine:latest

WORKDIR /app

COPY --from=build /build/voetbalpoules voetbalpoules

ENTRYPOINT ["/app/voetbalpoules"]