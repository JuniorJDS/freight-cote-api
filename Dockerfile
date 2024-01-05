FROM golang:1.21 as builder

RUN apt-get update && apt-get install -y \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local

RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx


WORKDIR /go/app
COPY . /go/app

# RUN go mod tidy
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

RUN strip --strip-unneeded main
RUN upx main


FROM alpine:3.7

# Create a non-root user called "appuser"
RUN adduser -D -g "" -s /bin/bash appuser

RUN apk add --no-cache tzdata

WORKDIR /go/app

EXPOSE $PORT

ENV TZ="UTC"

# Workaround to copy .env file just if it exists, useful for CD. And to copy

COPY --from=builder /go/app/main .

RUN chown -R appuser:appuser /go/app

# Set the user to "appuser"
USER appuser

CMD ["./main"]
