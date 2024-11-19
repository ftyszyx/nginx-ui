set CGO_ENABLED=1
go build -tags=jsoniter  -o nginx-ui.exe -v main.go
rem json库使用jsoniter