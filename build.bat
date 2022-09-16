go build -v -o main.exe -trimpath -ldflags "-s -w -buildid=" ./main

rem run as admin
main.exe -service install 
main.exe -service start