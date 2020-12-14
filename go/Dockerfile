FROM golang:1.15-alpine as BUILD
RUN apk add build-base
WORKDIR /testdir
COPY src .
RUN go build -o main .
RUN ["chmod", "+x", "scripts/runDockerTests.sh"]
ENTRYPOINT ["sh", "scripts/runDockerTests.sh"]
