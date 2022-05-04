#未成年运营社区平台
采用vue2+goframe开发
打包命令：

<!-- gf build cmd/cli/main.go -n=fed-cli -v=1.0 -a=amd64 -o=linux,darwin
gf build cmd/http/main.go -n=fed-http -v=1.0 -a=amd64 -o=linux,darwin -->
<!-- -n 指定编译后名称 -->
<!-- -v 指定编译后名称 -->
gf build cmd/cli/main.go -n=tpp-cli -v=1.0 -s=linux -a=amd64
gf build cmd/http/main.go -n=tpp-http -v=1.0 -s=linux -a=amd64


使用vendor包编译
gf build -mod=vendor cmd/cli/main.go -n=tpp-cli -v=1.0 -a=amd64 -o=linux,darwin
gf build -mod=vendor cmd/http/main.go -n=tpp-http -v=1.0 -a=amd64 -o=linux,darwin