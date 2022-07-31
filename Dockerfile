FROM golang:latest as builder

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . /go/src/app
RUN CGO_ENABLED=0 go build -o /go/bin/api ./main.go ./router.go ./middleware.go

FROM scratch as prod
COPY --from=builder /go/bin/api /go/bin/api
COPY --from=builder /go/src/app/config/config.yaml /config.yaml
CMD ["/go/bin/api"]
