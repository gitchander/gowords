package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	prmt "github.com/gitchander/permutation"
)

func main() {
	wordsFile()
	//wordsZip()
}

func wordsFile() {

	var (
		chars    string
		filename string
	)

	flag.StringVar(&chars, "chars", "", "characters set")
	flag.StringVar(&filename, "filename", "", "filename")

	flag.Parse()

	rs := []rune(chars)

	words, err := readWordsFromFile(filename)
	checkError(err)

	m := make(map[string]struct{})
	for _, word := range words {
		m[word] = struct{}{}
	}

	find(m, rs)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func find(words map[string]struct{}, rs []rune) {

	resultMap := make(map[string]struct{})

	total := 0
	for wordLen := 3; wordLen <= len(rs); wordLen++ {
		var count int
		f := func(word string) bool {
			if _, ok := words[word]; ok {
				if _, ok := resultMap[word]; !ok {
					resultMap[word] = struct{}{}
					fmt.Println(word)
					count++
					total++
				}
			}
			return true
		}
		testRunes(rs, wordLen, f)
		//fmt.Printf("words: len %d, count:%d\n", wordLen, count)
		//fmt.Println()
	}
	//fmt.Println("total count:", total)
}

func testRunes(rs []rune, k int, f func(word string) bool) {
	n := len(rs)
	ds := make([]rune, k)
	as := make([]int, k)
	for i := range as {
		as[i] = i
	}
	p := prmt.New(prmt.IntSlice(as))
	for {
		for p.Next() {
			for i, a := range as {
				ds[i] = rs[a]
			}
			if !f(string(ds)) {
				return
			}
		}
		overflow := nextComb(as, n)
		if overflow {
			break
		}
	}
}

func nextComb(as []int, n int) (overflow bool) {
	if (len(as) == 1) || nextComb(as[1:], n) {
		d := as[0] + 1
		if d > (n - len(as)) {
			return true
		}
		for i := range as {
			as[i] = d + i
		}
	}
	return false
}

func readWords(r io.Reader) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return words, nil
}

func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readWords(file)
}
