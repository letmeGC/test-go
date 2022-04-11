FROM golang:alpine
RUN mkdir /test
WORKDIR /test
COPY . .
RUN go build -o app  .
EXPOSE 8080
CMD ["/test/app"]