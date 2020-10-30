package main

import (
	"archive/zip"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func wordsZip() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("mast be argument")
	}
	var lang string
	var rs []rune
	if len(args) == 2 {
		lang = "en"
		rs = []rune(os.Args[1])
	} else {
		lang = args[1]
		rs = []rune(args[2])
	}

	words, err := wordsFromZipFile("words.zip", lang)
	checkError(err)

	m := make(map[string]struct{})
	for _, word := range words {
		m[word] = struct{}{}
	}

	find(m, rs)
}

func wordsFromZipFile(filename, lang string) (words []string, err error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, file := range r.File {
		if fileNameHasLang(file.Name, lang) {
			rc, err := file.Open()
			if err != nil {
				return words, err
			}
			defer rc.Close()
			ws, err := readWords(rc)
			if err != nil {
				return nil, err
			}
			words = append(words, ws...)
		}
	}
	return words, nil
}

func fileNameHasLang(fileName, lang string) bool {

	// format:
	// file-en.txt
	// lang: en

	s := fileName
	ext := filepath.Ext(s)
	if ext != "" {
		s = s[:len(s)-len(ext)]
	}
	i := strings.LastIndexByte(s, '-')
	if i == -1 {
		return false
	}
	return (s[i+1:] == lang)
}
