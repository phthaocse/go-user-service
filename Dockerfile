FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /user-service-go

EXPOSE 8000

CMD ["/user-service-go"]
