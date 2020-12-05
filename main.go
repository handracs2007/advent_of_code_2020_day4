package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "strconv"
    "strings"
)

func isValid(mf []string, validators map[string]func(string) bool, f []string) bool {
    for _, mField := range mf {
        exists := false
        valid := false

        for _, field := range f {
            if strings.HasPrefix(field, mField+":") {
                exists = true
                valid = validators[mField](field)
                break
            }
        }

        if !exists || (exists && !valid) {
            return false
        }
    }

    return true
}

func main() {
    fc, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Println(err)
        return
    }

    validators := make(map[string]func(string) bool)
    validators["byr"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]

        year, err := strconv.Atoi(val)
        return err == nil && year >= 1920 && year <= 2002
    }

    validators["iyr"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]

        year, err := strconv.Atoi(val)
        return err == nil && year >= 2010 && year <= 2020
    }

    validators["eyr"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]

        year, err := strconv.Atoi(val)
        return err == nil && year >= 2020 && year <= 2030
    }

    validators["hgt"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]

        if strings.HasSuffix(val, "cm") {
            val = strings.ReplaceAll(val, "cm", "")
            height, err := strconv.Atoi(val)
            return err == nil && height >= 150 && height <= 193
        }

        if strings.HasSuffix(val, "in") {
            val = strings.ReplaceAll(val, "in", "")
            height, err := strconv.Atoi(val)
            return err == nil && height >= 59 && height <= 76
        }

        return false
    }

    validators["hcl"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]
        regex, _ := regexp.Compile("^#[a-f0-9]{6}$")
        return regex.MatchString(val)
    }

    validators["ecl"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]
        colours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

        for _, c := range colours {
            if c == val {
                return true
            }
        }

        return false
    }

    validators["pid"] = func(s string) bool {
        data := strings.Split(s, ":")
        val := data[1]
        regex, _ := regexp.Compile("^[0-9]{9}$")
        return regex.MatchString(val)
    }

    mf := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

    lines := strings.Split(string(fc), "\n")
    cf := make([]string, 0)
    validCount := 0

    for idx, line := range lines {
        data := strings.Split(line, " ")
        cf = append(cf, data...)

        if len(line) == 0 || idx == len(lines)-1 {
            if isValid(mf, validators, cf) {
                validCount++
            }

            cf = make([]string, 0)
            continue
        }
    }

    fmt.Println(validCount)
}
