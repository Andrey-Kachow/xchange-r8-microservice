FROM golang:1.23.1

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o ./xchange-r8-microservice_exe ./cmd/server/

EXPOSE 8080

CMD [ "./xchange-r8-microservice_exe" ]