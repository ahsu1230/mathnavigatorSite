REM build app in production
forfiles /P dist /M * /C "cmd /c if @isdir==TRUE rmdir /S /Q @file"
npm run-script build

REM upload to Git (push to brach test_json!)
git add components/repos/json/*.json
git commit -m "json update"
git push origin test_json

REM upload to FileZilla
ftp -i -s:u.ftp
