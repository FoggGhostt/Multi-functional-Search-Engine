FROM golang:1.23.4

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

# Указываем переменную окружения (может быть перезаписана в docker-compose.yml)
ENV MONGO_URI=mongodb://my-mongo:27017/InvertIndex

# CMD ["./main"]
CMD ["tail", "-f", "/dev/null"]
