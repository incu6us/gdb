package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
    "sync"
)

type BackupConfig struct {
    S3Config           S3Config            `json:"s3"`
    MySQLBackupConfigs []MySQLBackupConfig `json:"mysql-configs"`
    LocalBackupConfigs []LocalBackupConfig `json:"local-configs"`
}

type S3Config struct {
    BucketEndpoint string `json:"bucket-endpoint"`
    Region         string `json:"region"`
    Bucket         string `json:"bucket"`
    AccessKey      string `json:"access-key"`
    SecretKey      string `json:"secret-key"`
}

type MySQLBackupConfig struct {
    Host     string   `json:"host"`
    Port     int      `json:"port"`
    DB       string   `json:"db"`
    User     string   `json:"user"`
    Password string   `json:"password"`
    Options  []string `json:"options"`
    S3Dir    string   `json:"s3-dir"`
}

type LocalBackupConfig struct {
    Path  string `json:"path"`
    S3Dir string `json:"s3-dir"`
}

var backupConfig = new(BackupConfig)
var once sync.Once

func GetConfig() *BackupConfig {
    return backupConfig
}

func ReadConfig(confFile string) {
    once.Do(func() {
        f, err := os.Open(confFile)
        if err != nil {
            log.Printf("Can't open file: %v", err)
        }

        defer f.Close()

        buf, err := ioutil.ReadAll(f)
        if err != nil {
            log.Printf("IO read error: %v", err)
        }

        if err := json.Unmarshal(buf, backupConfig); err != nil {
            log.Printf("TOML error: %v", err)
        }
    })
}
