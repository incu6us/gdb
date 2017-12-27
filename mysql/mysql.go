package mysql

import (
    "log"
    "strconv"
    "strings"

    "gdb/config"

    "github.com/incu6us/barkup"
    "github.com/influxdata/influxdb/cmd/influxd/backup"
)

func Backup(mysqlConfig config.MySQLBackupConfig, s3Config config.S3Config) {

    log.Printf("MYSQl: %v", mysqlConfig)

    mysql := &barkup.MySQL{
        Host:     mysqlConfig.Host,
        Port:     strconv.Itoa(mysqlConfig.Port),
        DB:       mysqlConfig.DB,
        User:     mysqlConfig.User,
        Password: mysqlConfig.Password,
        Options:  mysqlConfig.Options,
    }

    var s3 *barkup.S3
    if &s3Config != nil {
        s3 = &barkup.S3{
            Endpoint:     s3Config.Endpoint,
            Region:       s3Config.Region,
            Bucket:       s3Config.Region,
            AccessKey:    s3Config.AccessKey,
            ClientSecret: s3Config.SecretKey,
        }
    }
    dir := mysqlConfig.S3Dir
    if !strings.HasSuffix(dir, "/") {
        dir += "/"
    }
    result := mysql.Export()
    if result.Error != nil {
        log.Panicf("Get MySQL result error: %s", result.Error)
    }

    if err := result.To(dir, s3); err != nil {
        log.Panicf("S3 error: %s", err)
    }

}
