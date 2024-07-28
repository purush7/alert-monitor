# Dockerfile for Alert Monitor
FROM golang:latest

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-api .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o alert_monitor ./cmd

RUN cp alert_monitor /root/

WORKDIR /root/
# Copy migrations
COPY internal_ext/repository/migrations migrations

CMD ["./alert_monitor"]
