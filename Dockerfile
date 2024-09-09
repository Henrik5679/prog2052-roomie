#syntax=docker/dockerfile:1
#https://docs.docker.com/guides/language/golang/build-images/
FROM golang:1.23.0
LABEL maintainer="wwwilhel@stud.ntnu.no"

# Set up execution environment in container's GOPATH
WORKDIR /source

# Set environment variables
ENV PORT=8080

# Download Go modules
COPY go.mod ./
RUN go mod download

# Move source code to container
COPY presentation ./presentation/
COPY main.go      ./

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /program

# For some reason no need to expose port 8080
#EXPOSE $PORT

CMD ["/program"]
