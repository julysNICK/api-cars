FROM golang:alpine AS builder

Label maintainer="Julys <Julysmartins54@gmail.com>"

RUN apk update && apk add --no-cache git


WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .  

EXPOSE 8080

CMD ["./main"]