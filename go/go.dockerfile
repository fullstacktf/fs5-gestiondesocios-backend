FROM golang:1.15-alpine
RUN mkdir -p /go/src/fs5-gestiondesocios-backend/src 
WORKDIR /go/src/fs5-gestiondesocios-backend/src
ADD ./src/ ./
RUN cd ../../
RUN export GOPATH=$(pwd) 
RUN apk add git
RUN apk upgrade
RUN apk add build-base
RUN go get github.com/gorilla/mux
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/mysql
RUN go get github.com/steinfletcher/apitest
RUN go get -u github.com/basgys/goxml2json
RUN go get -u github.com/tidwall/gjson 
RUN go get github.com/rs/cors
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
EXPOSE 8080
CMD ["go","run","main.go"]

