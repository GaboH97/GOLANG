package main 

import (
   // "code.google.com/p/go-tour/wc"
    "strings"
    "fmt"
    "math"
)

func ContadorPalabras(s string) map[string]int {
    m := make(map[string]int)
    strs := strings.Fields(s)
    for i := 0; i< len(strs);i++{
    	v, ok := m[strs[i]]
        if  v ==0 && !ok{
            m[strs[i]] = 1
        }else{
            m[strs[i]] = m[strs[i]] +1
        }
    }
    
    return m
}

func fibonacci() func() int {
    sum :=0
    current :=1
    
    return func () int{
    	sum  = sum + current
        current = sum-current
        return current
    }
    
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
    
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
        fmt.Println(math.Tan(math.Pow10(i)))
    }
    
    word := "I am what I am"
    fmt.Println(word)
    m := ContadorPalabras(word)
    for k :=  range m{
    	fmt.Println("key[%s] value[%s]\n", k, m[k])
    }
    
    fmt.Println(Reverse("hola"))
}






