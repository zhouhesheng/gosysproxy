go build -v -o main.exe -trimpath -ldflags "-s -w -buildid=" ./main
main.exe -service install 
main.exe -service start