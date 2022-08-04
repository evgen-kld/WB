package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type param struct {
	f []int
	d string
	s bool

	files []string
}

func getParams() *param {
	f := flag.String("f", "", "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	d := flag.String("d", "\t", "use DELIM instead of TAB for field delimite")
	s := flag.Bool("s", false, "do not print lines not containing delimiters")

	flag.Parse()
	tmp := strings.Split(*f, ",")

	fields := make([]int, len(tmp))

	for i := range tmp {
		num, err := strconv.Atoi(tmp[i])
		if err != nil || num == 0 {
			return nil
		}
		fields[i] = num
	}

	args := &param{
		f: fields,
		d: *d,
		s: *s,
	}

	args.files = append(args.files, flag.Args()...)

	return args
}

func readFile(filename string) []string {
	var lines []string
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// cutLines - cut lines by specified args
func cutLines(p *param, lines []string) []string {
	delimiter := "\t"

	if p.d != delimiter {
		delimiter = p.d
	}

	var result []string

	for _, line := range lines {
		if delimiter != "" && strings.Contains(line, delimiter) {
			words := strings.Split(line, delimiter)

			cutLine := strings.Builder{}

			for _, val := range p.f {
				if len(words) >= val {
					cutLine.WriteString(words[val-1])
					cutLine.WriteString(delimiter)
				}
			}

			// trim extra delimiter
			result = append(result, strings.TrimSuffix(cutLine.String(), delimiter))

		} else if !p.s {
			result = append(result, line)
		}
	}

	return result
}

func cut() []string {

	var result []string

	args := getParams()

	for _, val := range args.files {
		lines := readFile(val)

		cutLines := cutLines(args, lines)
		result = append(result, cutLines...)
	}

	return result
}

func main() {
	lines := cut()
	for i := range lines {
		fmt.Println(lines[i])
	}
}
