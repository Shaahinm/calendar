# syntax=docker/dockerfile:1

FROM golang:1.23.0

#only to test and for sqlite, should be removed
RUN apt update
RUN apt install -y build-essential libsqlite3-dev

WORKDIR /app

#COPY go.mod go.sum ./
COPY . .
RUN go mod download
#RUN pwd

# Build
WORKDIR /app/cmd
#RUN ls -a
#CGO_ENABLED=1 only to test and for sqlite, should be removed
#RUN CGO_ENABLED=1 GOOS=linux go build -o /docker-calendar
RUN GOOS=linux go build -o /docker-calendar
WORKDIR /app
#EXPOSE 8000

# Run
CMD ["/docker-calendar"]
#docker build -t go-calendar --progress plain .