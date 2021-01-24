FROM golang:1.15

WORKDIR /home/app

ADD . .

CMD ["make", "run"]