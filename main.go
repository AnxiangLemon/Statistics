package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("本程序用于统计字符串出现的次数 -- by 银时")
		fmt.Println("linux|win使用方式 : (./Statistics)|(Statistics.exe) <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	counts := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Convert the map to a slice of entries
	entries := make([]struct {
		line  string
		count int
	}, 0, len(counts))
	for line, count := range counts {
		entries = append(entries, struct {
			line  string
			count int
		}{line, count})
	}

	// Sort the slice by count
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].count < entries[j].count
	})

	// Print the sorted entries
	for _, entry := range entries {
		fmt.Printf("%s (%d 次)\n", entry.line, entry.count)
	}
	fmt.Println("字符串数量 ", len(counts))
}
