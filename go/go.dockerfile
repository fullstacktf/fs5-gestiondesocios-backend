FROM golang:1.15-alpine
WORKDIR /app
COPY /src/ /app
RUN go mod download

RUN go build -o /out/example .
#RUN go run server.go
ENTRYPOINT ["tail", "-f", "/dev/null"]
