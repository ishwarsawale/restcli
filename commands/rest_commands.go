package commands

import (
  "github.com/urfave/cli"
  "restcli/rest"
  "fmt"
  "os"
  "net"
  "time"
)

var (
    SERVER_FLAG      string
    SERVER_PORT    string
  	CACERT_FLAG    string
    CERT_FLAG      string
    KEY_FLAG      string
    DATA_FLAG      string
)


var REST_Commands = []cli.Command{
   {
        Name:  "get",
        Usage: "get",
        Action: getAction,
      },

   {
        Name:  "post",
        Usage: "post",
        Action: postAction,
      },




}


func getAction(c *cli.Context) {
  SERVER_FLAG = c.GlobalString("server")
  SERVER_PORT = c.GlobalString("port")
  CACERT_FLAG = c.GlobalString("cacert")
  CERT_FLAG   = c.GlobalString("cert")
  KEY_FLAG   = c.GlobalString("key")

  // check blank args
  if SERVER_FLAG =="" || KEY_FLAG == "" || CACERT_FLAG == "" || CERT_FLAG == "" {
    cli.ShowAppHelp(c)
    os.Exit(1)
  }

  //check server ping
  seconds := 5
  timeOut := time.Duration(seconds) * time.Second
  _, err := net.DialTimeout("tcp", SERVER_FLAG+":"+SERVER_PORT, timeOut)
  if err != nil {
    fmt.Println("Invalid argument Detail: connect failed in tcp_connect()")
     os.Exit(1)
  }



  if c.Args().First() !="" {

   rest.Get_Rest(SERVER_FLAG,SERVER_PORT,CACERT_FLAG,CERT_FLAG,KEY_FLAG,c.Args().First())
  }else{
    fmt.Println("Invalid argument Detail: no end point added")
  }
}


//post action

func postAction(c *cli.Context) {
  SERVER_FLAG = c.GlobalString("server")
  SERVER_PORT = c.GlobalString("port")
  CACERT_FLAG = c.GlobalString("cacert")
  CERT_FLAG   = c.GlobalString("cert")
  KEY_FLAG   = c.GlobalString("key")
  DATA_FLAG   = c.GlobalString("data")
  
// check blank args
  if SERVER_FLAG =="" || KEY_FLAG == "" || CACERT_FLAG == "" || CERT_FLAG == "" {
    cli.ShowAppHelp(c)
    os.Exit(1)
  }

  //check server ping
  seconds := 5
  timeOut := time.Duration(seconds) * time.Second
  _, err := net.DialTimeout("tcp", SERVER_FLAG+":"+SERVER_PORT, timeOut)
  if err != nil {
    fmt.Println("Invalid argument Detail: connect failed in tcp_connect()")
     os.Exit(1)
  }


  //call if endpoint is given
  if c.Args().First() !="" {
   rest.Post_Rest(SERVER_FLAG,SERVER_PORT,CACERT_FLAG,CERT_FLAG,KEY_FLAG,c.Args().First(),DATA_FLAG)
  }else{
    fmt.Println("Invalid argument Detail: no end point added")
  }
}
