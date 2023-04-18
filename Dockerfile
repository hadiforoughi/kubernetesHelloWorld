FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /kubernetesHelloWorld

EXPOSE 8080

CMD ["/kubernetesHelloWorld"]