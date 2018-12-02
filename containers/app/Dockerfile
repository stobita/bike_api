FROM golang:1.11

WORKDIR /bike_api
COPY . /bike_api
RUN go get -u github.com/oxequa/realize
CMD ["realize", "start", "--run", "--no-config"] 
