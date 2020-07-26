FROM golang:1.11.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

# We want to build our application's binary executable
RUN go get github.com/gorilla/websocket
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
# We have to copy the output from our
# builder stage to our production stage
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app ./app

# we can then kick off our newly compiled
# binary exectuable!!
CMD ["./app/main"]