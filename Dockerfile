FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY server/go.mod server/go.sum ./
RUN go mod download && go mod verify

COPY server/. .
RUN go build -v -o /usr/local/bin/app ./...
EXPOSE 5000
EXPOSE 8080
CMD ["app"]