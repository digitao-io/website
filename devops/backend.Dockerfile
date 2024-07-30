FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY ./backend ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o ./backend

###

FROM alpine AS release-stage

WORKDIR /app

COPY --from=build-stage /app/backend ./
COPY ./devops/backend.prod.json ./

RUN adduser -D nonroot
USER nonroot
ENV SERVICE_CONFIG_PATH=./backend.prod.json

EXPOSE 3000

ENTRYPOINT ["./backend"]
