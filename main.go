package main
import (
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"os/user"
	"github.com/atotto/clipboard"
)
func main(){
	var pkg=flag.String("f","","")
	flag.Parse()
	var user,err=user.Current()
	if err!=nil {
		fmt.Println(err)
	}
	if *pkg!="" {
		var fileContent,err=ioutil.ReadFile(user.HomeDir+"/.gopkglst")
		if err!=nil {
			fmt.Println(err)
		}
		var s=string(fileContent)
		var lst=strings.Split(s,"\n")
		for i:=0; i<len(lst); i++ {
			var tmp=strings.Split(lst[i]," ")
			if tmp[0]==*pkg {
				fmt.Println(tmp[1])
				clipboard.WriteAll(tmp[1])
			}
		}
	}
}