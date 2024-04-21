FROM golang:1.22.0

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apt update \
    && apt clean \
    && rm -r /var/lib/apt/lists/*

RUN apt install git && \
    apt install curl

COPY go.mod go.sum* ./

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]