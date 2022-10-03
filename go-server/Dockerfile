FROM golang:alpine AS build

WORKDIR /app
ADD ./src/ /app/
RUN go build -o server

FROM alpine
WORKDIR /app
COPY --from=build /app /app/
ENTRYPOINT [ "./server" ]
