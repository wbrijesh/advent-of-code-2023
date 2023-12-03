package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
    numbers_string := formatString(lines[i])

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

    fmt.Println(number)

    total = total + number
  }

  fmt.Println(total)
}

func formatString(inputString string) string {
  replacedString := inputString

  one_index := strings.Index(inputString, "one")
  two_index := strings.Index(inputString, "two")
  three_index := strings.Index(inputString, "three")
  four_index := strings.Index(inputString, "four")
  five_index := strings.Index(inputString, "five")
  six_index := strings.Index(inputString, "six")
  seven_index := strings.Index(inputString, "seven")
  eight_index := strings.Index(inputString, "eight")
  nine_index := strings.Index(inputString, "nine")

  sorted_index_values := []int{one_index, two_index, three_index, four_index, five_index, six_index, seven_index, eight_index, nine_index}
  sort.Sort(sort.IntSlice(sorted_index_values))
  index_values := []int{one_index, two_index, three_index, four_index, five_index, six_index, seven_index, eight_index, nine_index}

  occurance_sequence := []string{}

  for i := 0; i < len(index_values); i++ {
    if sorted_index_values[i] != -1 {
      if sorted_index_values[i] == one_index {
	occurance_sequence = append(occurance_sequence, "one")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == two_index {
	occurance_sequence = append(occurance_sequence, "two")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == three_index {
	occurance_sequence = append(occurance_sequence, "three")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == four_index {
	occurance_sequence = append(occurance_sequence, "four")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == five_index {
	occurance_sequence = append(occurance_sequence, "five")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == six_index {
	occurance_sequence = append(occurance_sequence, "six")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == seven_index {
	occurance_sequence = append(occurance_sequence, "seven")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == eight_index {
	occurance_sequence = append(occurance_sequence, "eight")
	sorted_index_values[i] = -1
      } else if sorted_index_values[i] == nine_index {
	occurance_sequence = append(occurance_sequence, "nine")
	sorted_index_values[i] = -1
      }
    }
  }

  for i := 0; i < len(occurance_sequence); i++ {
    if occurance_sequence[i] == "one" {
      replacedString = strings.ReplaceAll(replacedString, "one", "1e")
    } else if occurance_sequence[i] == "two" {
      replacedString = strings.ReplaceAll(replacedString, "two", "2o")
    } else if occurance_sequence[i] == "three" {
      replacedString = strings.ReplaceAll(replacedString, "three", "3e")
    } else if occurance_sequence[i] == "four" {
      replacedString = strings.ReplaceAll(replacedString, "four", "4")
    } else if occurance_sequence[i] == "five" {
      replacedString = strings.ReplaceAll(replacedString, "five", "5e")
    } else if occurance_sequence[i] == "six" {
      replacedString = strings.ReplaceAll(replacedString, "six", "6")
    } else if occurance_sequence[i] == "seven" {
      replacedString = strings.ReplaceAll(replacedString, "seven", "7")
    } else if occurance_sequence[i] == "eight" {
      replacedString = strings.ReplaceAll(replacedString, "eight", "8t")
    } else if occurance_sequence[i] == "nine" {
      replacedString = strings.ReplaceAll(replacedString, "nine", "9e")
    }
  }

  replacedString = regexp.MustCompile(`[^0-9]`).ReplaceAllString(replacedString, "")

  return replacedString
}

