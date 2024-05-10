FROM golang:1.21

WORKDIR /usr/src/app

RUN go install github.com/codegangsta/gin@latest
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
EXPOSE 3001
ENTRYPOINT ["gin","run","main.go"]