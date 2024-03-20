
`sudo tar -C /usr/local -zxvf go$VERSION.$OS-$ARCH.tar.gz`    


`sudo vim /etc/profile`   (for a system-wide installation)  

```
export MYPARAM=go1-15-11

export GOROOT=/usr/local/$MYPARAM
export GOPATH=${HOME}/${MYPARAM}-path
export GOPROXY=https://goproxy.io,direct
export GOPRIVATE=gitlab.com,gitee.com,gitea.com

export PATH=$PATH:${GOROOT}/bin:${GOPATH}/bin
```

`source /etc/profile`  



`source ~/.bashrc`  