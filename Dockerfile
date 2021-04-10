FROM golang:1.15-alpine as builder

ENV CGO_ENABLED=0

# if you use docker buildx, it doesnt needed
#ENV GOOS=linux
#ENV GOARCH=arm64
WORKDIR /go/notify_hatebu_reaction

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add tzdata

COPY --from=builder /go/notify_hatebu_reaction/app /app
CMD ["/app"]