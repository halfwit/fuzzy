package main

import( 
	"flag"
	"fmt"
	"os"

    "github.com/google/gofuzz"
)

var times = flag.Int("n", 1000, "Number of outputs")
var length = flag.Int("b", 100, "Length of each line for output")

func main() {
	flag.Parse()

	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(0)
	}

	f := fuzz.New()
	for i := 0; i < *times; i++ {
		b := make([]byte, *length)
		f.Fuzz(&b)
		fmt.Printf("%s\n", b)
	}
}