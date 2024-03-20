
`protoc -I=$SRC_DIR --go_out=$DST_DIR              $SRC_DIR/addressbook.proto`  

`protoc -I=$SRC_DIR                                         --go_out=plugins=grpc:$DST_DIR $SRC_DIR/addressbook.proto`  
`protoc -I D:\software\development\go-path-1-19-4\src -I .\ --go_out=plugins=grpc:.        proto\abc.proto`  

