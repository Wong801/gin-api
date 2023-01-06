# On Golang Will build the projects 
FROM golang:1.19-alpine as build

WORKDIR /app 

COPY . ./

RUN go mod tidy 

RUN go build -o gin-api main.go 

# Stage Deployer
# OS Alpine Will Running The app with copy all data from build stage
FROM alpine as main

WORKDIR /app

COPY --from=build /app/gin-api /app 

CMD ["./gin-api"]

EXPOSE 8000