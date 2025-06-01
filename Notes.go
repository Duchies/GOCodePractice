point 1) 

fmt.Sprintf("%d",10) 
// it return string // you can convert int to string,
but string to int is not possible with this methods
you need to use stringconv.Atoi() to do so.


point 2)
stringBuilder in golang 
func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

point 3)

i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)

# the strings package
s := "hello world"
strings.Contains(s, "wo")   // true
strings.Split(s, " ")       // ["hello", "world"]
strings.ToUpper(s)          // "HELLO WORLD"
strings.ToLower(s)          // "hello world"
strings.TrimSpace("  hi ")  // "hi"
strings.ReplaceAll(s, "l", "L") // "heLLo worLd"
	
