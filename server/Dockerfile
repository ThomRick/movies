FROM golang:latest 
RUN mkdir /app 
ADD . /go/src/github.com/mouminoux/movies/server
WORKDIR /go/src/github.com/mouminoux/movies/server
RUN go get
RUN go build
CMD ["server"]
