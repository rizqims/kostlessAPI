FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o kostless
ENTRYPOINT ["/app/kostless-api"]
