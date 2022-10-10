FROM golang:1.19.2 AS builder

WORKDIR /go/src/github.com/seanmalloy/gsd
COPY . .
RUN make build

FROM scratch

COPY --from=builder /go/src/github.com/seanmalloy/gsd/gsd /gsd

ENTRYPOINT ["/gsd"]