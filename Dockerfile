FROM golang:alpine
RUN apk --no-cache add ca-certificates

COPY . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify
RUN go mod tidy
RUN go mod vendor

COPY . .
EXPOSE 5000
ENV GO111MODULE=on
ENV GOOS=linux
RUN go build -ldflags="-s -w" -o tracks_api cmd/service/main.go

CMD [ "/app/tracks_api" ]