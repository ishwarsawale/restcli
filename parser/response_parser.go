package parser
import (
  "github.com/buger/jsonparser"
  "fmt"
)

func Parse_json( response []byte ) {
        certname,_,_,_ :=  jsonparser.Get(response,"[0]")
				fmt.Println(string(certname))
}
