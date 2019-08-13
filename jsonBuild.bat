forfiles /P dist /M * /C "cmd /c if @isdir==TRUE rmdir /S /Q @file"
npm run-script build
