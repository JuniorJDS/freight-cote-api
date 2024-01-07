FROM golang:1.21.5

# Create a non-root user called "appuser"
RUN useradd -m -s /bin/bash appuser

WORKDIR /go/app

COPY go.* ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -v -o series-api-go .

RUN chown -R appuser:appuser /go

# Set the user to "appuser"
USER appuser

EXPOSE 5000
