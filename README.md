# ms-curso-auth-grpc

# Catalog Mod
go mod init github.com/j-ew-s/ms-curso-auth-grpc

## Porta
5500

# Comunica  com outros serviços:

1 - Catalog API (:5200)
2 - Cart API (:5300)
3 - Ordering API (:5400)

# PROTO : Gerar serviços GRPC 
Cria  auth.pb.go   em auth-services
 protoc --go_out=plugins=grpc:auth-services  auth.proto


# DOCKER 

docker build --tag auth-grpc:v0.0.1 -f Dockerfile   .

docker compose build

docker compose up
