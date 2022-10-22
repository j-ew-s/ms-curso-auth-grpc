# ms-curso-auth-grpc



# PROTO : Gerar serviços GRPC 
Cria  auth.pb.go   em auth
protoc --go_out=plugins=grpc:auth  user.proto
