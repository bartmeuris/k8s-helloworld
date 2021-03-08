FROM golang:latest as builder

COPY main.go /main.go

RUN CGO_ENABLED=0 go build -o /helloworld /main.go

FROM alpine:latest

COPY --from=builder /helloworld /helloworld

ENTRYPOINT [ "/helloworld" ]