FROM alpine:latest

RUN apk add --no-cache git make musl-dev go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go build -o api-playstore

CMD ["./api-playstore"]