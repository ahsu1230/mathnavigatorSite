rm dist/*
npm run-script build

# upload to Git (push to brach test_json!)
git add components/repos/json/*.json
git commit -m "json update"
git push origin test_json

# upload to FileZilla
