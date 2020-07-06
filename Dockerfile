FROM golang:1.13-alpine3.11

RUN mkdir /g2ww

ADD . /g2ww

WORKDIR /g2ww

RUN go build -o g2ww .

ENV DOCKER=tsuki
CMD ["/g2ww/g2ww"]