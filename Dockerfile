FROM golang:1.23.5
RUN mkdir /app
COPY . /app
# RUN go mod download
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o shoedesk
EXPOSE 8080
CMD ["/app/shoedesk"]
