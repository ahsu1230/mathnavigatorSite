del ftp.tmp

echo open %MATH_NAV_HOST% >> ftp.tmp
echo %MATH_NAV_USER% >> ftp.tmp
echo %MATH_NAV_PW% >> ftp.tmp

echo binary >> ftp.tmp
echo cd public_html >> ftp.tmp
echo lcd dist >> ftp.tmp
echo mput * >> ftp.tmp

echo disconnect >> ftp.tmp
echo quit >> ftp.tmp