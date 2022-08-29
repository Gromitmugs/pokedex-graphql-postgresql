FROM golang:1.19

WORKDIR $GOPATH/src/github.com/Gromitmugs/pokedex-graphql-posgresql

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /pokedex-docker

EXPOSE 8001

CMD [ "/pokedex-docker" ]

