FROM golang:1.15.2-alpine

RUN apk update && \
  apk add git

RUN mkdir /app
WORKDIR /app

# realize
# https://qiita.com/rin1208/items/64a6bc469d19ad0ec981
RUN go get -u github.com/oxequa/realize
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GO111MODULE=on
EXPOSE 8080
CMD ["realize", "start", "--build", "--run"]
