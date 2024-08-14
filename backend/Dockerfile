# Specifies a parent image
FROM golang:1.22-bookworm

# Creates an app directory to hold your app’s source code
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* .

# Installs Go dependencies
RUN go mod download

## Copies everything from your root directory into /app
COPY . .

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD ["air", "-c", ".air.toml"]