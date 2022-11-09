FROM golang:1.18-alpine AS build

WORKDIR /build
COPY ./ .

RUN go mod download

RUN CGO_ENABLED=0 go build -o voetbalpoules

FROM alpine:latest

WORKDIR /app

COPY --from=build /build/voetbalpoules voetbalpoules

ENTRYPOINT ["/app/voetbalpoules"]