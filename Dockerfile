FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

FROM scratch
COPY --from=0 /app /app
ENTRYPOINT ["/app"]
