package main

import (
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/sundeck-io/CommentQL/go/pkg/qtag"
	"io"
	"os"
)

func main() {
	usage := flag.Bool("usage", false, "whether or not to show usage")
	all := flag.Bool("all", false, "include all unknown comments and attributes (default false)")
	arbitraryAttributes := flag.Bool("allattributes", false, "include all unknown attributes for known comments (default false)")
	flag.Parse()

	if *usage {
		flag.Usage()
		return
	}

	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	input := string(stdin)

	e := qtag.NewExtractor()
	tags := e.Extract(input, *all, *arbitraryAttributes)

	str, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
