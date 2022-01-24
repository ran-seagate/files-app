FROM golang:alpine

RUN mkdir "files-app"
WORKDIR /files-app

RUN export GO111MODULE=on
RUN apk update && apk add git
RUN git clone https://github.com/ran-seagate/files-app.git
WORKDIR /files-app/files-app
RUN go build .
RUN chmod -R 755 .

EXPOSE 8081

ENTRYPOINT [ "./files-app" ]