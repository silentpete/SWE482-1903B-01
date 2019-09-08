# build the golang binary
FROM golang:1.13-alpine3.10 as builder

WORKDIR /go/src/bundys

COPY . .

RUN \
  go get -v -d ./... \
  && env GOOS=linux GOARCH=386 go build -v


FROM scratch

COPY --from=builder /go/src/bundys/bundys /bin/

EXPOSE 6060

ENTRYPOINT ["/bin/bundys"]
