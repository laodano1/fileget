
# --plugin=protoc-gen-go=..\initialize\tools\protoc-gen-go.exe
protoc --proto_path=. --micro_out=. --go_out=. gameprovide.proto
