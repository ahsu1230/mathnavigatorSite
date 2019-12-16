To update json files,

1.
Look at Google Excel Sheet
https://docs.google.com/spreadsheets/d/1SCxykC95iVLIiC1qX2nWyYcxRyEEtCs6zqkv9cOzifE/edit?usp=sharing

Request edit/write access from administrator.


2.
Look at each of the sheets and follow their formats to edit!
* purple columns mean primary_keys!
* blue columns mean auto-generated!


3.
Once finished editing,
File > Download as...

and select .csv
Only one .csv file is created for every page, so you will have to export
each Google sheet page as a .csv file.


4.
Go to: https://www.csvjson.com/csv2json
Upload a .csv file
and press CONVERT!


5.
On the right side, you should see a json be generated.
Press DOWNLOAD to download the .json file.


6.
Make sure the file is named appropriately and
replace the old __.json file with the new one.


7.
Test the site locally... (run `./testJson.sh`)
If things look incorrect, edit the Google Sheets and repeat steps 3-6.

Once finished,
Run `./buildJson.sh`
and you're finished!
