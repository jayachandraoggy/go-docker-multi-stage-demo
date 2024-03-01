FROM golang:alpine
LABEL authors="jayachandra"

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o goapp .
EXPOSE 3000

CMD [ "/app/goapp" ]