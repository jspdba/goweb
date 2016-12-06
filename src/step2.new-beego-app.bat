@echo 设置 App 的值为您的应用文件夹名称
set APP=web
set GOPATH=%~dp0..
set BEE=%GOPATH%\bin\bee
%BEE% new %APP%
cd %APP%
echo %BEE% run %APP% > run.bat
echo pause >> run.bat
start run.bat
pause
start http://127.0.0.1:8888