package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  data, err := os.ReadFile("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  lines := strings.Split(string(data), "\n")

  total := 0

  for i := 0; i < len(lines) - 1; i++{ 
    numbers_string := regexp.MustCompile(`[^0-9]`).ReplaceAllString(lines[i], "")
    number := 0

    if len(numbers_string) == 1 {
      numbers_string = numbers_string + numbers_string
      number, err = strconv.Atoi(numbers_string)
      if err != nil {
	fmt.Println(err)
      }
    } 

    if len(numbers_string) == 2 {
      number, err = strconv.Atoi(numbers_string)
      if err != nil {
	fmt.Println(err)
      }
    }

    if len(numbers_string) > 2 {
      numbers_string = numbers_string[0:1] + numbers_string[len(numbers_string)-1:]
      number, err = strconv.Atoi(numbers_string)
      if err != nil {
	fmt.Println(err)
      }
    }

    total = total + number
  }

  fmt.Println(total)
}
