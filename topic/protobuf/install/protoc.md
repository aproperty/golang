
`sudo apt-get install autoconf automake libtool curl make g++ unzip`  

`git clone https://github.com/protocolbuffers/protobuf.git`  
`cd protobuf`  
`git submodule update --init --recursive`  
`./autogen.sh`  

`./configure`  
`make`  
`make check`  
`sudo make install`  
`sudo ldconfig` # refresh shared library cache.  

`protoc --version`  




或者在下载页直接下载：  
https://github.com/protocolbuffers/protobuf/releases

