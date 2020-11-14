FROM golang:1.15.3-buster
#RUN apt-get update -y
WORKDIR /app
COPY /src/ /app
RUN go build -o /out/example .
