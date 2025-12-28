FROM golang:1.22

WORKDIR /app

RUN apt-get update \
 && apt-get install -y --no-install-recommends poppler-utils \
 && rm -rf /var/lib/apt/lists/*

COPY go.mod ./
RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/wpscli"]
