FROM golang:1.15 as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN groupadd -g 10001 hackz-api \
  && useradd -u 10001 -g hackz-api hackz-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app ./cmd/api

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd

EXPOSE 8080

USER hackz-api

ENTRYPOINT ["/go/bin/app"]