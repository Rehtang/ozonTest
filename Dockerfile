FROM golang:1.20.5

ARG port=9090
ARG storage="postgres"
ARG postgres-link="jdbc:postgresql://localhost:5432/"


WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install -y git

RUN go mod download

RUN go build -o main .

CMD ["./main"]
