#Remove old messages
rm -force proto/messages.pb.go

#Generate from prot file
protoc `
 -I="D:\Source\src\github.com\gogo\protobuf\protobuf" `
 -I="D:\Source\src\github.com\AsynkronIT\protoactor-contracts" `
 -I="D:\Source\src" `
 -I="proto" `
 --gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mactor.proto=github.com/AsynkronIT/protoactor-go/actor,plugins=grpc:proto `
 proto/messages.proto

# #Fix double import
# #get-content proto/messages.pb.go | select-string -pattern 'import google_protobuf1 "github.com/gogo/protobuf/types"\nimport google_protobuf2 "github.com/gogo/protobuf/types"`|import google_protobuf "github.com/gogo/protobuf/types' -notmatch | Out-File proto/messages.pb.go
# (Get-Content proto/messages.pb.go).replace('google_protobuf1', 'google_protobuf') | Set-Content proto/messages.pb.go
# (Get-Content proto/messages.pb.go).replace('google_protobuf2', 'google_protobuf') | Set-Content proto/messages.pb.go
