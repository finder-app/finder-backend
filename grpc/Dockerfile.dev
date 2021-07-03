FROM golang:1.15.2-alpine

WORKDIR /app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git \
  && go get github.com/gravityblast/fresh

COPY . .

CMD ["fresh"]
