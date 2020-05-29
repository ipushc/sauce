# build stage
FROM ipushc/golangxnode:1.13.8-v12 AS builder

LABEL stage=sauce-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/sauce

RUN cd /go/src/sauce && CGO_ENABLED=0 go build -mod vendor


# final stage
FROM alpine:3.10.2

COPY --from=builder /go/src/sauce/sauce ./

CMD ["./sauce"]
