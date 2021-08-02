# golang builder
FROM golang:latest AS builder
WORKDIR /go/src/github.com/hugdubois/svc-fizzbuzz/
COPY . .
#RUN apk add --no-cache --update git make protobuf protobuf-dev ca-certificates curl && \
     #rm -rf /var/cache/apk/*
RUN rm -f /go/src/github.com/hugdubois/svc-fizzbuzz/_build/svc-fizzbuzz
RUN make

# minimal image from scratch
FROM scratch
LABEL maintainer="Hugues Dubois <hugdubois@gmail.com>"
#COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/github.com/hugdubois/svc-fizzbuzz/_build/svc-fizzbuzz /svc-fizzbuzz
EXPOSE 8080
ENTRYPOINT ["/svc-fizzbuzz"]
CMD ["serve"]
