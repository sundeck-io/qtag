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
	unidentified := flag.Bool("unidentified", false, "whether or not to include unidentified comment patterns")
	unexpected := flag.Bool("unexpected", false, "whether to include unexpected attributes of known of comment patterns")
	flag.Parse()

	if *usage {
		flag.Usage()
		return
	}

	flag.Bool("empty", true, "whether or not to include empty comment prefixes")
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	input := string(stdin)

	e := qtag.NewExtractor()
	tags := e.Extract(input, *unidentified, *unexpected)

	str, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
