FROM golang:1.23  AS build

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . /api/

ENV CGO_ENABLED=0 GOOS=linux

RUN  go build -o /rest-api  ./main.go


FROM alpine:latest

COPY --from=build /rest-api .

EXPOSE 3000

CMD [ "./rest-api" ]
