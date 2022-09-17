go build -v -o main.exe -trimpath -ldflags "-s -w -buildid=" ./main

rem run as admin
rem main.exe -service install 
rem main.exe -service start
rem main.exe -service stop
rem main.exe -service uninstall 
