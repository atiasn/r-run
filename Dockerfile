
FROM golang:alpine as builder
ARG software=uploadToAli
ARG GOOS=linux
ARG GOARCH=amd64
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Set necessary environmet variables needed for our image
ENV GOPROXY=https://goproxy.io \
    CGO_ENABLED=0 \
    GOOS=$GOOS \
    GOARCH=$GOARCH

RUN go build -o $software .

FROM alpine
ARG software=uploadToAli
COPY --from=builder /build/$software /$software
COPY ./cacert.pem /etc/ssl/certs/
CMD ["/uploadToAli"]
