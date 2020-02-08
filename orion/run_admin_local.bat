:: This script builds the website locally AND runs the orion server locally
:: So you may interact with the website and api server end to end.

:: First step is to make sure mysql is installed on your computer and running
:: net start MySQL
:: net stop MySQL
:: mysql db --user=root --password

:: Second step is to build the website locally
cd sites\admin
call npm run build-local
cd ..\..

:: Third step is to run the api server locally
go run main.go configs\config_local.yaml

:: Test by going to http://localhost:8080 in an internet browser
