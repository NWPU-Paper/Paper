go get github.com/astaxie/beego
go get github.com/beego/bee

cd %GOPATH%\src
go build github.com\beego\bee
copy bee.exe %GOPATH%\bin\bee.exe
del bee.exe
