



In order to generate Go code, 
the Go package's import path        must be provided for every .proto file 
(including those transitively depended upon by the .proto files being generated).     

 



The Go import path is locally specified in a .proto file 
by declaring a `go_package` option 
with 
the full import path 
of the Go package.   






Example usage:  

```
option go_package = "example.com/project/protos/fizz";
```