FROM golang:1.20.3-bullseye

WORKDIR /app

COPY . .

RUN echo "PORT=4040\nDB_HOST=db\nDB_USER=root\nDB_PASS=root" > .env

EXPOSE 4040
CMD ["go", "run", "main.go"]