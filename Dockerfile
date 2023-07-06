FROM golang:1.20.5

WORKDIR /app

COPY . .

ENV PORT=9090
ENV STORAGE_TYPE=postgres
ENV POSTGRES_LINK=postgres://postgres:postgres@localhost:5432/ozonTest

RUN go build -o ozonTest .

CMD ["./ozonTest", "-port=${PORT}", "-storage=${STORAGE_TYPE}", "-postgres-link=${POSTGRES_LINK}"]
