package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"time"
)

// This CLI sends http requests to a LOCAL orion webserver.
// This script should not be used for a non-local orion webserver!
// Make sure orion is healthy before running this CLI
//
// How to use this CLI
// In the orion folder, run `go test -v =count=1 ./... > orion.test.log`
// While will produce a file that contains the full output of the test
//
// Then traverse to this folder (or open a new console tab/window) and run these commands.
// go run render.go
// go run render.go --filePath ../orion/orion.test.log
// go run render.go --filePath samples/orion.test.compile.log
//
// by default the program will pick up the file at `orion/orion.test.log`.
// but you can specify any file using the `filePath` flag

var DEFAULT_FILE_PATH string = "../orion/orion.test.log"

func main() {
	log.Println("Starting render program...")

	var filePath string
	flag.StringVar(&filePath, "filePath", DEFAULT_FILE_PATH, "Path to test log")
	flag.Parse()

	// Read incoming test.log file
	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open input file (%v)", err)
	}
	defer inputFile.Close()

	// Create index.html
	outputFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Could not create file index.html (%v)", err)
	}
	defer outputFile.Close()

	writeToOutputFile(inputFile, outputFile)
	log.Println("Finished writing index.html")
}

// This function reads the *.test.log file to collect data
// And writes to a temporary *.html file
// We then take a second pass by readin the temp html file
// And insert a "header" to the real index.html file
func writeToOutputFile(inputFile, outputFile *os.File) {
	tmpFilePath := "index.tmp.html"
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		log.Fatalf("Could not create file temporary index.tmp (%v)", err)
	}
	defer tmpFile.Close()

	writer := bufio.NewWriter(tmpFile)
	writer.WriteString("<html>\n")
	writer.WriteString("<head><link rel=\"stylesheet\" href=\"style.css\"/></head>\n")
	writer.WriteString("<body>\n")
	writer.WriteString("<h1>Test Results</h1>\n")

	var hasCompileErrors = false
	var numTests  = 0
	var failedTestNames []string
	var currentTestName string

	// Read inputFile and write to outputFile as we go
	var builder strings.Builder
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasSuffix(line, "[build failed]") {
			writer.WriteString(createCompileErrorMessage(line))
			hasCompileErrors = true
			break
		}

		if strings.HasPrefix(line, "=== RUN") {
			numTests++
			currentTestName = extractCurrentTestName(line)
			builder.Reset()
			builder.WriteString("<b>" + line + "</b><br/>")
		} else if strings.HasPrefix(line, "--- PASS:") {
			writer.WriteString("<p class=\"success\">\n")
			writer.WriteString(builder.String())
			writer.WriteString(line)
			writer.WriteString("</p>\n")
			continue
		} else if strings.HasPrefix(line, "--- FAIL:") {
			writer.WriteString("<p class=\"failed\">\n")
			writer.WriteString(builder.String())
			writer.WriteString(line)
			writer.WriteString("</p>\n")
			failedTestNames = append(failedTestNames, currentTestName)
			continue
		} else if strings.HasPrefix(line, "time=") {
			builder.WriteString("<span class=\"log\">" + line + "</span><br/>")
		} else if strings.HasPrefix(line, "[GIN]") {
			builder.WriteString("<span class=\"gin\">" + line + "</span><br/>")
		} else {
			builder.WriteString(line + "<br/>")
		}
	}
	writer.WriteString("</body></html>\n")
	writer.Flush()
	log.Println("Finish first write to temporary file.")

	// Begin second pass (reset file descriptor)
	time.Sleep(time.Duration(2) * time.Second)
	tmpFile.Seek(0, 0)

	writer2 := bufio.NewWriter(outputFile)
	scanner2 := bufio.NewScanner(tmpFile)
	for scanner2.Scan() {
		line := scanner2.Text()
		writer2.WriteString(line)
		if strings.Contains(line, "<h1>Test Results</h1>") {
			writeHeader(writer2, hasCompileErrors, numTests, failedTestNames)
		}
	}
	writer2.Flush()
	log.Println("Finish second write (with header) to output file. Deleting temporary file.")
	os.Remove(tmpFilePath)
}

func writeHeader(writer *bufio.Writer, hasCompileErrors bool, numTests int, failedTestNames []string) {
	if hasCompileErrors {
		writer.WriteString("<h3>Some tests cannot run until code can successfully compile. " + 
			"Please resolve the compilation errors.</h3>\n")
		return
	}

	writer.WriteString(fmt.Sprintf("<h3>%d Tests were run</h3>\n", numTests))

	if !hasCompileErrors && len(failedTestNames) == 0 {
		writer.WriteString("<h3 class=\"success\">All Tests successfully passed!</h3>\n")
	}

	if len(failedTestNames) > 0 {
		writer.WriteString(fmt.Sprintf(
			"<div class=\"failed\">\n" + 
			"<h3>Failed %d Tests</h3>\n" + 
			"<ul>\n", 
			len(failedTestNames)))
		for _, testName := range failedTestNames {
			writer.WriteString(fmt.Sprintf("<li>%s</li>\n", testName))
		}
		writer.WriteString("</ul></div>\n")
	}	
}

func createCompileErrorMessage(line string) string {
	return fmt.Sprintf(
		"<h2 class=\"error compile\">There's a compile error!</h2>\n" + 
		"<p>%s</p>\n",
		line)
}

// Assumes the line looks like "=== RUN TestXYZ"
func extractCurrentTestName(line string) string {
	elems := strings.Split(line, " ")
	return elems[4]
}