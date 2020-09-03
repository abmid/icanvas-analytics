FROM golang:1.14

# Add Maintainer Info
LABEL maintainer="Abdul Hamid (abdul.surel@gmail.com)"

# Set Workdir
WORKDIR /go/src/app

# COPY all files to image
COPY . .

# Install Nodejs
RUN apt-get update \
    && apt-get -y install curl gnupg \ 
    && curl -sL https://deb.nodesource.com/setup_12.x  | bash - \
    && apt-get -y install nodejs

# Build production web client
RUN cd web/app && npm run build

# Install Dependency
RUN go get -d -v ./...
RUN go install -v ./...

# Build Golang
RUN cd cmd/app && go build -o main .

# Expose port
EXPOSE 8181

WORKDIR /go/src/app/cmd/app

CMD ["./main"]