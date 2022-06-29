# Base image, tagged "build"
FROM golang:1.18.3 as build

# Specify the source code directory (creates it)
WORKDIR /src/app
# We need to create this directory to alude to it
RUN mkdir /bin/app

# ENV DB_HOST localhost
# ENV DB_PORT 5432

# Copy files and install dependencies
COPY . /src/app
# Apparently, these two are not copied with the prior command (investigate)
COPY go.mod  go.sum ./
# Downloads (and installs(?)) dependencies specified in go.mod (investigate)
RUN go mod download

# Build into bin directory
# -o tag indicates where to build into, next argument where to take packages
RUN go build -o /bin/app ./cmd/todolist


# Image to work (distroless, low weight)
FROM gcr.io/distroless/base-debian11

# Copy files from "build" to distroless image (only the compiled binaries so the image does not weight a lot)
COPY --from=build /bin/app /
COPY ./web/static ./web/static
COPY ./swagger.yml ./swagger.yml
 
# Port to listen to
EXPOSE 8000

# Start the app
CMD ["/todolist"]