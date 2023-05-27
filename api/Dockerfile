FROM golang:1.20.4-alpine3.18

RUN mkdir /api
WORKDIR /api

# Copying the manifests
COPY go.mod go.sum ./

# Downloading dependencies
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o main .

# Exposing the port
EXPOSE 8081

# Entry point
CMD [ "./main" ]