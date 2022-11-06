
`protoc -I=$SRC_DIR --go_out=$DST_DIR              $SRC_DIR/addressbook.proto`  

`protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR $SRC_DIR/addressbook.proto`  
