# Compile stage
FROM golang:1.14 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /heyemoji

# Final stage
FROM debian:buster

WORKDIR /
COPY --from=build-env /heyemoji /

CMD ["/heyemoji"]
