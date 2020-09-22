echo "Running orion tests..."
go test -v -count=1 ./... > orion.test.log

echo "Rendering test results..."
go run ../clients/golangTestRenderer/render.go -filePath "../orion/orion.test.log"
mv index.html ../clients/golangTestRenderer/index.html
open -a "Google Chrome" ../clients/golangTestRenderer/index.html