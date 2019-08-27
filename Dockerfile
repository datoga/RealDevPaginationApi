FROM golang AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/PaginationApi
WORKDIR /go/src/PaginationApi

ADD . /go/src/PaginationApi

WORKDIR /go/src/PaginationApi

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -ldflags '-w -s -extldflags "-static"' .

FROM alpine

RUN mkdir /app 
WORKDIR /app
COPY --from=builder /go/src/PaginationApi/RealDevPaginationApi .

COPY db.json .

EXPOSE 8080

CMD ["/app/RealDevPaginationApi"]
