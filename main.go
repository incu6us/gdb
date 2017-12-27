package main

import (
    "flag"
    "log"
    "os"

    "gdb/config"
    "gdb/mysql"
)

var (
    confFile string
    cmdHelp  bool
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    flag.StringVar(&confFile, "conf", "config.json", "Choose a config file")
    flag.BoolVar(&cmdHelp, "h", false, "Help")
    flag.Parse()
    if cmdHelp {
        flag.PrintDefaults()
        os.Exit(1)
    }

    config.ReadConfig(confFile)

    s3Config := config.GetConfig().S3Config

    for _, mysqlConfig := range config.GetConfig().MySQLBackupConfigs {
        mysql.Backup(mysqlConfig, s3Config)
    }
}
