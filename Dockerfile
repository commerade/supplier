FROM golang:1.20 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 1323
ENTRYPOINT ["./server"]
CMD $CMD_PARAM