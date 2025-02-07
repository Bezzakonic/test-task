package main

import (
 "bufio"
 "fmt"
 "os"
 "regexp"
 "strconv"
 "strings"
)

type Str struct {
 A        string
 Operator string
 B        string
 Numstr   string
}

func main() {
 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Введите выражение: ")
 var input string
 scanner.Scan()
 input = scanner.Text()
 re := regexp.MustCompile(`^"([^"]+)"\s*([+\-*\/])\s*("([^"]+)"|(\d+))$`)
 matches := re.FindStringSubmatch(input)
 if matches == nil {
  panic("Ошибка: Неверный формат ввода.")
 }
 var s Str
 s.A = matches[1]
 s.Operator = matches[2]
 s.B = matches[4]
 s.Numstr = matches[5]

 ress := ""

 switch s.Operator {
 case "+":
  ress = s.A + s.B
 case "-":
  ress = s.Strminus()
 case "*":
  num, err := strconv.Atoi(s.Numstr)
  if err != nil || num < 1 || num > 10 {
   panic("Число должно быть от 1 до 10 включительно.")
  }
  ress = s.Umnostr(num)
 case "/":
  num, err := strconv.Atoi(s.Numstr)
  if err != nil || num < 1 || num > 10 {
   panic("Число должно быть от 1 до 10 включительно.")
  }
  ress = s.Delstr(num)
 default:
  panic("Недопустимый оператор. Используйте +, -, *, /.")
 }

 if len(ress) > 40 {
  ress = ress[:40] + "..."
 }

 fmt.Printf("%q\n", ress)
}

func (s *Str) Strminus() string {
 return strings.ReplaceAll(s.A, s.B, "")
}

func (s *Str) Umnostr(numstr int) string {
 return strings.Repeat(s.A, numstr)
}

func (s *Str) Delstr(numstr int) string {
 leght := len(s.A)
 if numstr <= 0 || leght < numstr {
  return ""
 }
 return s.A[:leght/numstr]
}
