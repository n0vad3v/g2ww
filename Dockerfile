FROM golang:stretch AS build

COPY .  /go/src
WORKDIR /go/src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/g2ww *.go

FROM alpine

COPY --from=build /go/bin/g2ww /g2ww
ENV DOCKER=tsuki
CMD ["/g2ww"]