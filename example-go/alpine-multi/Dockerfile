FROM golang:1.13.5
WORKDIR /go/src/github.com/narenaryan/app
COPY code/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /go/src/github.com/narenaryan/app .
CMD ["./app"]