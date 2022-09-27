FROM golang:1.19.1-alpine3.15

RUN apk update && apk add --no-cache git

CMD sh
