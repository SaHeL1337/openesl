package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Template struct {
	Width      int64  `json:"width"`
	Height     int64  `json:"height"`
	Background []int  `json:"background"`
	Text       []Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
	Color   []int  `json:"color"`
	Size    int    `json:"size"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}

func parseTemplate() (Template, error) {
	var template Template
	templateFile, err := os.Open("template.json")
	if err != nil {
		return Template{}, fmt.Errorf("failed to open template.json: %v", err)
	}
	defer templateFile.Close()

	tbd, err := io.ReadAll(templateFile)
	if err != nil {
		return Template{}, fmt.Errorf("failed to read template.json: %v", err)
	}
	err = json.Unmarshal(tbd, &template)
	if err != nil {
		return Template{}, fmt.Errorf("failed to unmarshal template.json: %v", err)
	}
	return template, nil
}
