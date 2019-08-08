package cf

import (
	"ch2/temconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}
		f := temconv.Fahrenheit(t)
		c := temconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, temconv.FToC(f), c, temconv.CToF(c))

	}

}
