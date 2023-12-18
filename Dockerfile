FROM golang:latest

WORKDIR /app


COPY . /Desafio-CLA
CMD ["tail", "-f", "/dev/null"]