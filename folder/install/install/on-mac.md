 
`tar -C /usr/local -zxvf go$VERSION.$OS-$ARCH.tar.gz`     



`vi ~/.zshrc`   

```sh
export MYTEMP=go1-20-1
export GOROOT=/usr/local/$MYTEMP
export GOPATH=${HOME}/${MYTEMP}-path
export GOSUMDB=off
export GOPROXY=https://goproxy.io,direct
export GOPRIVATE=gitlab.com,gitee.com,gitea.com

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# GOSUMDB="sum.golang.org"
# export GOPROXY=https://goproxy.cn,direct
```

`source ~/.zshrc`   
