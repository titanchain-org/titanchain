FROM golang:1.21-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /titanchain
RUN cd /titanchain && make titan

FROM alpine:latest

WORKDIR /titanchain

COPY --from=builder /titanchain/build/bin/titan /usr/local/bin/titan

RUN chmod +x /usr/local/bin/titan

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/titan"]

CMD ["--help"]
