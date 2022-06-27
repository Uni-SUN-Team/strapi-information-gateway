FROM golang:1.18.3

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN rm -f /usr/src/app/src/resource/app.yaml && mv /usr/src/app/.resource/sit/app.yaml /usr/src/app/src/resource/app.yaml
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]