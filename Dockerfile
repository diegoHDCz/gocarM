FROM golang:1.22.4-alpine

WORKDIR /autdoor

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /autdoor/cmd/autdoor

RUN go build -o /autdoor/bin/autdoor .

EXPOSE 8080
ENTRYPOINT [ "/autdoor/bin/autdoor" ]