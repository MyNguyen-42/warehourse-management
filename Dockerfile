FROM golang:1.18.4-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./
ENTRYPOINT ["/out/main"]
EXPOSE 8080

#APP_ENV=prod go run main.go