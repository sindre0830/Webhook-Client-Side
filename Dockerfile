FROM golang:1.16

LABEL maintainer "rickarl@stud.ntnu.no"

# Create a new directory for our code
RUN mkdir /app
ADD . /app
WORKDIR /app

# Get go.mod and main.go files
ADD ./go.mod /
ADD ./main.go /

# Download dependencies and build - put our executable in its own folder for gitignore
RUN CGO_ENABLED=0 GOOS=linux go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./bin/main

# Ports we want to expose for this service
EXPOSE 8081 8081

# Give permissions to main and run
RUN chmod +x ./bin/main

CMD ["./bin/main"]
