FROM golang:1.15.3-buster
RUN apt-get update
WORKDIR /app
COPY /src/go/ /app
RUN go build -o /out/example .
