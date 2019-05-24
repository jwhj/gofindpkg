package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	var pkg = flag.String("f", "", "")
	var isCurrent = flag.Bool("c", false, "")
	flag.Parse()
	var user, err = user.Current()
	if err != nil {
		fmt.Println(err)
	}
	if *pkg != "" {
		var fileContent, err = ioutil.ReadFile(user.HomeDir + "/.gopkglst")
		if err != nil {
			fmt.Println(err)
		}
		var s = string(fileContent)
		var lst = strings.Split(s, "\n")
		for i := 0; i < len(lst); i++ {
			var tmp = strings.Split(lst[i], " ")
			if tmp[0] == *pkg {
				fmt.Println(tmp[1])
				clipboard.WriteAll(tmp[1])
			}
		}
	} else if *isCurrent {
		var path, _ = os.Getwd()
		var gopaths = strings.Split(os.Getenv("GOPATH"), ":")
		for i := 0; i < len(gopaths); i++ {
			if strings.HasPrefix(path, gopaths[i]) {
				var s = strings.Replace(path, gopaths[i]+"/src/", "", -1)
				fmt.Println(s)
				clipboard.WriteAll(s)
			}
		}
	}
}
