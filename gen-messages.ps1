#Remove old messages
rm -force messages.pb.go

#Generate from prot file
protoc `
 -I="D:\Source\src\github.com\gogo\protobuf\protobuf" `
 -I="D:\Source\src\github.com\AsynkronIT\protoactor-contracts" `
 -I="D:\Source\src" `
 -I="proto" `
 -I="." `
 --gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mactor.proto=github.com/AsynkronIT/protoactor-go/actor,plugins=grpc:. `
 proto/messages.proto