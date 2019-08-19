git add components/repos/json/*.json
git commit -m "json update"
git push origin test_json

createFtp.bat
ftp -i -s:ftp.tmp
