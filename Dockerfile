FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOFLAGS=-mod=vendor
ENV APP_USER app

RUN mkdir /capacitacion
ADD . /capacitacion
WORKDIR /capacitacion

RUN go mod vendor
RUN go build . 

EXPOSE 8000
CMD ["/capacitacion/capacaction"]