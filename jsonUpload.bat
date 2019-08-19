git add components/repos/json/*.json
git commit -m "json update"
git push origin test_json

echo creating ftp file...
call createFtp.bat

echo upload to ftp...
ftp -i -s:ftp.tmp
