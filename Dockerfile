# Compile stage
FROM golang:1.17 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /heyemoji

# Final stage
FROM debian:buster

WORKDIR /

# Copy app executable from builder container
COPY --from=build-env /heyemoji /

# Copy CA certificates to prevent x509: certificate signed by unknown authority errors
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/heyemoji"]
