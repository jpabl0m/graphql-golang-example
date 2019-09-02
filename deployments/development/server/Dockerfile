FROM golang:1.12.9-alpine3.9
RUN apk add --update --no-cache \
      git \
      gcc \
      libc-dev
COPY . /graphql
WORKDIR /graphql
RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u -t github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
RUN go get github.com/cortesi/modd/cmd/modd
CMD ["modd", "-f", "configs/modd.conf"]
