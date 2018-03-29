FROM golang:latest AS builder

RUN mkdir -p /go/src/github.com/cjimti/go-echo
COPY . /go/src/github.com/cjimti/go-echo

RUN go get ...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/tcp-echo ./src/github.com/cjimti/go-echo

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/tcp-echo /tcp-echo

WORKDIR /

ENTRYPOINT ["/tcp-echo"]
