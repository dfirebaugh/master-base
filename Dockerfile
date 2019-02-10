# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info 
LABEL maintainer="Dustin Firebaugh <dafirebaugh@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dfirebaugh/master-base

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# build main.go
RUN go build -o master-base src/*.go

# Run the executable
CMD ["./master-base"]
