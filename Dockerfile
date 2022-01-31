FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o example .

FROM alpine:3.9.5
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/example .

CMD [ "./example" ]