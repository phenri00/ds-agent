FROM golang:1.8 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ds-agent

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/ds-agent .
CMD ["./ds-agent"]  
 
