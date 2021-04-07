FROM alpine:3.12 AS build

RUN apk --no-cache add go

ADD . /go/src/github.com/safecornerscoffee/golang-crud-api
WORKDIR /go/src/github.com/safecornerscoffee/golang-crud-api

RUN go build

FROM alpine:3.12

COPY --from=build /go/src/github.com/safecornerscoffee/golang-crud-api/golang-crud-api /usr/bin/golang-crud-api

CMD ["/usr/bin/golang-crud-api"]

