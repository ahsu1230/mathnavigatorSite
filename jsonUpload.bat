git add components/repos/fetchers/json/*.json
git commit -m "json update"
git push origin master

echo creating ftp file...
call createFtp.bat

echo upload to ftp...
ftp -i -s:ftp.tmp
