package main

import (
    "flag"
    //"github.com/alecthomas/chroma/v2"
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
    //"reflect"
)
func show_uncommented_lines(filePath string) {

    var matchingLines string

    file, err := os.Open(filePath)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    r, err := regexp.Compile("^[[:space:]]*#")

    if err != nil {
        log.Fatal(err)
    }

    for scanner.Scan() {
        if ! r.MatchString(scanner.Text()) {
            matchingLines = fmt.Sprintf("%s\n%s", matchingLines, scanner.Text())
        }
    }

    fmt.Println(matchingLines)
} 

func bulk_grepper(fpaths []string) {

    
    for i := 0; i < len(fpaths); i++ {
    	current_fpath := fpaths[i]
    	show_uncommented_lines(current_fpath)
    }
}

func main() {

    var regex string
    flag.StringVar(&regex, "regex", "^[[:space:]]*#", "file to search!")

    flag.Parse()

    fpaths := flag.Args()
    bulk_grepper(fpaths)
}


