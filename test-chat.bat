@echo off
echo Opening Terminal Chat Test Environment...
echo.

echo Starting server in new window...
start "Chat Server" cmd /k "go run server/main.go"

echo Waiting 2 seconds for server to start...
timeout /t 2 >nul

echo Starting first client...
start "Chat Client 1" cmd /k "go run client/main.go"

echo Waiting 1 second...
timeout /t 1 >nul

echo Starting second client...
start "Chat Client 2" cmd /k "go run client/main.go"

echo.
echo Chat test environment started!
echo - Server window: "Chat Server" 
echo - Client windows: "Chat Client 1" and "Chat Client 2"
echo.
pause 