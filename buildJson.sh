rm dist/*
parcel build index.html

# upload to Git (dangerous to push to master!)
git add js/repos/json/*.json
git commit -m "json update"
git push origin change_json

# upload to FileZilla
