FROM golang

RUN mkdir "files-app"
WORKDIR /files-app

RUN export GO111MODULE=on
COPY . .
RUN chmod -R 755 .
RUN go build

EXPOSE 8081

ENTRYPOINT [ "./files-app" ]