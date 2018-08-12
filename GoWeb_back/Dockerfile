# Get latest golang docker image.
FROM golang:latest

# Create a directory inside the container to store our web-app and then make it working directory.
RUN mkdir -p /go/src/web-app
WORKDIR /go/src/web-app

# Copy the web-app directory into the container.
COPY . /go/src/web-app

# Download and install third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable
ENV PORT 8080

# Expose port 8080 to the host so that outer-world can access your application
EXPOSE 8080

# Tell Docker what command to run when the container starts
CMD ["go-wrapper", "run"]