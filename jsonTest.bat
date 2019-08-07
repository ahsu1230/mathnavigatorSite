set target=C:\Users\Amylin\Downloads\jsons\*
dir %target%

@echo off
set count=0
for %%x in (%target%.json) do set /a count+=1
echo %count%

IF %count% GTR 0 (
	echo files available!
	move %target%.json components\repos\json
)ELSE (
	echo no files!
)

npm run-script run
