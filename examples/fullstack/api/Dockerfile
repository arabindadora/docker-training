FROM golang:alpine as builder
WORKDIR /go/src/api
COPY src .
RUN CGO_ENABLED=0 go build -o /go/bin/api -ldflags '-extldflags "-static"'

FROM scratch
COPY --from=builder /go/bin/api /api
EXPOSE 8000
ENTRYPOINT ["/api"]
