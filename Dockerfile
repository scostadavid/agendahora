# Use a base image with Node.js and npm installed for building frontend assets
FROM node:16 AS node-builder
WORKDIR /app
COPY . .
RUN npm install -g tailwindcss
RUN npx tailwindcss -i cmd/web/assets/css/input.css -o cmd/web/assets/css/output.css

# Use the Templ image to generate templates
FROM ghcr.io/a-h/templ:latest AS templ-builder
WORKDIR /app
COPY . .
RUN ["templ", "generate"]

# Build the Go application
FROM golang:latest AS go-builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api/main.go

CMD ["./main"]

# Create a minimal image for the Go application
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /app
# COPY --from=go-builder main main
# EXPOSE 8080
# RUN ls

