package main

import (
    "flag"
    "log"
    "os"

    "github.com/incu6us/gdb/config"
    "github.com/incu6us/gdb/local"
    "github.com/incu6us/gdb/mysql"
)

var (
    confFile string
    cmdHelp  bool
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    flag.StringVar(&confFile, "conf", "", "Choose a config file")
    flag.BoolVar(&cmdHelp, "h", false, "Help")
    flag.Parse()
    if confFile == "" || cmdHelp {
        flag.PrintDefaults()
        os.Exit(1)
    }

    config.ReadConfig(confFile)

    s3Config := config.GetConfig().S3Config

    for _, mysqlConfig := range config.GetConfig().MySQLBackupConfigs {
        mysql.Backup(mysqlConfig, s3Config)
    }

    for _, localConfig := range config.GetConfig().LocalBackupConfigs {
        local.Backup(localConfig, s3Config)
    }
}
