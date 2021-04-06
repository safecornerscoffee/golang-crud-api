FROM alpine:3.12 AS build

RUN apk --no-cache add go

ADD . /go/src/github.com/safecornerscoffee/echo-mvc
WORKDIR /go/src/github.com/safecornerscoffee/echo-mvc

RUN go build

FROM alpine:3.12

COPY --from=build /go/src/github.com/safecornerscoffee/echo-mvc/echo-mvc /usr/bin/echo-mvc

CMD ["/usr/bin/echo-mvc"]

