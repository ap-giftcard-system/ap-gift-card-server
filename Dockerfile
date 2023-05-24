# Build stage
## Pull golang image from the hub
FROM golang:alpine AS builder

## Set up ENV vars
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## Choose work directory
WORKDIR /ap-gift-card-server

## Copy local project to docker container
COPY . .

## Run build command
RUN go build -o ap-gift-card .

# Run stage
## alpine:latest image is a light linux image
FROM alpine:latest AS runner

## Choose work directory
WORKDIR /ap-gift-card-server

## Copy the executable binary file and .env file from the last stage to the new stage
COPY --from=builder /ap-gift-card-server/ap-gift-card .

# Execute the build
CMD ["./ap-gift-card"]