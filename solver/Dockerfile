FROM golang:1.17-alpine as builder

WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/myapp

FROM scratch
COPY --from=builder /bin/myapp .
EXPOSE 8080

ENTRYPOINT [ "./myapp" ]