FROM golang:1.22.0

WORKDIR /usr/src/app

COPY . .

EXPOSE 8080

# Run
CMD ["go", "run", "."]