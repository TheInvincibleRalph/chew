package docx

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"

	"github.com/mmatongo/chew/internal/utils"
)

func ProcessDocx(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	var contents []string

	for _, file := range zipReader.File {
		if file.Name == "word/document.xml" {
			contents, err = utils.ExtractTextFromXML(file)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	var allContent strings.Builder
	for _, content := range contents {
		allContent.WriteString(content)
		allContent.WriteString(" ")
	}

	return []string{allContent.String()}, nil

	/*
		// In the event we just want chunks we can just return contents
		return contents, nil
	*/
}
