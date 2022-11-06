

ETCD：https://github.com/etcd-io/etcd





```sh
./etcd 
--listen-client-urls 'http://0.0.0.0:2379'    
--advertise-client-urls 'http://0.0.0.0:2379'  


# 后台启动 
# nohup etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379' &
```

