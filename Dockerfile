FROM golang:alpine AS build

# Add Maintainer Info
LABEL maintainer="guillaume.penaud@gmail.com"

RUN \
  apk add --no-cache git &&\
  mkdir /application

ADD . /application/
WORKDIR /application

# Download all the dependencies
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux \
  go build \
    -a -installsuffix cgo \
    -o /needys-output-producer \
  /application/cmd/needys-output-producer-listener/main.go

# ---------------------------------------------------------------------------- #

FROM alpine:latest

RUN adduser --system --disabled-password --home /needys-output-producer needys-output-producer

WORKDIR /needys-output-producer
USER needys-output-producer

COPY --from=build /needys-output-producer .

EXPOSE 8010

CMD ["./needys-output-producer"]
