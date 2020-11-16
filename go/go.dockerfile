FROM golang:1.15-alpine
RUN mkdir -p /go/src/fs5-gestiondesocios-backend/src 
WORKDIR /go/src/fs5-gestiondesocios-backend/src
ADD ./src/ ./
RUN cd ../../
RUN export GOPATH=$(pwd) 
RUN apk add git
RUN apk upgrade
RUN ls -lh
RUN go get github.com/gorilla/mux
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/mysql
EXPOSE 8080
CMD ["go","run","main.go"]

