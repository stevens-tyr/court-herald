FROM golang:1.11.2
WORKDIR /app
COPY . /app
RUN go build -o court_herald
ENV GIN_MODE release
CMD ["./court_herald"]