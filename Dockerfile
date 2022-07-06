# Base image, tagged "build"
FROM golang:1.18.3 as build

# Specify the source code directory (creates it)
WORKDIR /src/app

# We need to create this directory to alude to it
RUN mkdir /bin/app

# Copy files and install dependencies
COPY . /src/app
COPY go.mod  go.sum ./

# Downloads and installs dependencies specified in go.mod 
RUN go mod download

# Build into bin directory
# -o tag indicates where to build into, next argument where to take packages
RUN go build -o /bin/app ./cmd/todolist


# Image to work (distroless, low weight)
<<<<<<< HEAD
# FROM gcr.io/distroless/base-debian11

# Image to work, accepts command line 
FROM ubuntu

=======
FROM gcr.io/distroless/base-debian11

# Image for tests; comment the previous line and uncomment this one
# FROM ubuntu

>>>>>>> 22143b9fcb3cea715c4724ec5ce018e5b85b6682
# Copy files from "build" to distroless image (only the binaries so the image does not weight a lot)
COPY --from=build /bin/app /
COPY ./web/static ./web/static
COPY ./swagger.yml ./swagger.yml
 
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# Port to listen to
EXPOSE 8000