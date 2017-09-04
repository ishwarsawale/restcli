package commands

import (
  "github.com/urfave/cli"
)

var (
    PUPPET_SERVER      string
    PUPPET_PORT        string
    PUPPET_CACERT      string
    PUPPET_CERT        string
    PUPPET_KEY         string
    PUPPET_DATA         string
)


var Puppet_Flags =  []cli.Flag {
    cli.StringFlag{
      Name: "server, s",
      Usage: "full qualified hostname or ip add",
      Destination:&PUPPET_SERVER,
    },

    cli.StringFlag{
      Name: "port, p",
      // Value: "english",
      Usage: "port",
      Destination:&PUPPET_PORT,
    },

    cli.StringFlag{
      Name: "cacert, ca",
      // Value: "english",
      Usage: "cacert file",
      Destination:&PUPPET_CACERT,

    },

    cli.StringFlag{
      Name: "cert, ce",
      // Value: "english",
      Usage: "certificate",
      Destination:&PUPPET_CERT,
    },

    cli.StringFlag{
      Name: "key, ke",
      // Value: "english",
      Usage: "key",
      Destination:&PUPPET_KEY,
    },
    
    cli.StringFlag{
      Name: "data, da",
      // Value: "english",
      Usage: "data",
      Destination:&PUPPET_DATA,
    },

  }

func AuthFlags() (string, string,string, string, string) {
  server:= PUPPET_SERVER
  port:= PUPPET_PORT
  cacert:= PUPPET_CACERT
  cert:= PUPPET_CERT
  key:= PUPPET_KEY
return server,port,cacert,cert,key
}

