FROM golang:latest
LABEL maintainer='Adam Liang'
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o task_scheduler .
EXPOSE 8080
CMD ["./task_scheduler"]
