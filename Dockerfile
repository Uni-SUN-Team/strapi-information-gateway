FROM golang:1.17.9

ENV NODE=production
ENV PORT=8080
ENV CONTEXT_PATH=/strapi-information-gateway
ENV HOST_STRAPI=http://cms-cms-2:1337
ENV TOKEN_STRAPI=6102511ec9d936311659e59bd511783c8a1f3b4b00f47b27fbc16df3a4d2594924189cfa0412d435966fc188b6573e7a378323537e3d67b3cdbe87a0791012aa8425bd11f70266afa897ec5cb1b82dbb57529da17f273d94ec336b93c0a5cf9b056f1d52ba6d9576685227508c4736f8d06a9158d37ea665fd07f5a81bb6e72e
ENV SENTRY_DNS=https://60831186464d48a19f8160f0e885d01c@o1238561.ingest.sentry.io/6390163

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]