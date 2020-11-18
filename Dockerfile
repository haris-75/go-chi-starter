FROM golang:1.15.2-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /usr/local/my-project/
COPY . .

RUN go build -o /go/bin/my-project

FROM scratch

COPY --from=builder /go/bin/my-project /go/bin/my-project

EXPOSE 6969

FROM builder

ENTRYPOINT ["/go/bin/my-project"]
