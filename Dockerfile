FROM golang:1.19.4-alpine3.17 AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

WORKDIR $GOPATH/src/mypackage/myapp/cmd/server

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/studio-sol-back-end-test

FROM scratch

COPY --from=builder /go/bin/studio-sol-back-end-test /go/bin/studio-sol-back-end-test

# changed from CMD to ENTRYPOINT to make it work in heroku
# see this: https://devcenter.heroku.com/articles/container-registry-and-runtime#dockerfile-commands-and-runtime
ENTRYPOINT ["/go/bin/studio-sol-back-end-test"]