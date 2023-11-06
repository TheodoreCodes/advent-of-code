package skeleton

import (
	"advent-of-code/util"
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var fs embed.FS

const (
	inputUrl = "https://adventofcode.com/%d/day/%d/input"
)

func Run(year int, sessionCookie string) {
	log.Printf("Starting generating skeleton for year %d \r\n", year)
	if sessionCookie == "" {
		log.Fatalf("session cookie is mandatory")
	}

	for day := 1; day <= 25; day++ {
		log.Printf("Creating day number %d \n", day)
		path := fmt.Sprintf(
			"%s\\..\\..\\%d\\day%02d",
			util.Dirname(),
			year,
			day,
		)

		tmplData := map[string]any{
			"Day": fmt.Sprintf("%02d", day),
		}

		createDir(path)
		createFile("main.go", path, tmplData)
		createFile("main_test.go", path, nil)

		inputFile := createFile("input.txt", path, nil)
		downloadInput(inputFile, sessionCookie, year, day)
	}
}

func createDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("error creating directories: %v", err)
	}
}

func createFile(filename, filepath string, data map[string]any) *os.File {
	log.Printf("Creating file %s\r\n", filename)

	_, err := os.Stat(fmt.Sprintf("%s/%s", filepath, filename))
	if err == nil {
		log.Fatalf("filepath already exists: %s", filepath)
	}

	ts, err := template.ParseFS(fs, "templates/"+filename)
	if err != nil {
		log.Fatalf("error parsing filename template %s: %v", filename, err)
	}

	f, err := os.Create(
		fmt.Sprintf(
			"%s/%s",
			filepath, filename,
		))
	if err != nil {
		log.Fatalf("error creating filename %s/%s: %v", filepath, filename, err)
	}

	if err := ts.ExecuteTemplate(f, filename, data); err != nil {
		log.Fatalf("error executing template %s: %v", filename)
	}

	return f
}

func downloadInput(f *os.File, sessionCookie string, year, day int) {
	log.Printf("Creating input file")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(inputUrl, year, day), nil)
	if err != nil {
		log.Fatalf("error creating GET request: %v", err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error requesting input: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status while fetching input: %d", resp.StatusCode)
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("error writing input file: %v")
	}
}
