FROM golang:1.17.3  AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o  /ms-curso-auth-grpc



FROM gcr.io/distroless/base-debian10 

WORKDIR /
COPY --from=build /ms-curso-auth-grpc  /auth-grc

EXPOSE 5500

USER nonroot:nonroot

CMD ["/auth-grpc"]