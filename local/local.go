package local

import (
    "log"
    "strings"

    "github.com/incu6us/gdb/config"

    "github.com/incu6us/barkup"
)

func Backup(localConfig config.LocalBackupConfig, s3Config config.S3Config) {

    log.Printf("Local --> path: %s; s3: %s", localConfig.Path, localConfig.S3Dir)

    location := &barkup.Location{
        Path: localConfig.Path,
    }

    var s3 *barkup.S3
    if &s3Config != nil {
        s3 = &barkup.S3{
            BucketEndpoint: s3Config.BucketEndpoint,
            Region:         s3Config.Region,
            Bucket:         s3Config.Region,
            AccessKey:      s3Config.AccessKey,
            ClientSecret:   s3Config.SecretKey,
        }
    }

    dir := localConfig.S3Dir
    if !strings.HasSuffix(dir, "/") {
        dir += "/"
    }

    result := location.Export()
    if result.Error != nil {
        log.Printf("Get location result error: %s", result.Error)
    }

    if err := result.To(dir, s3); err != nil {
        log.Printf("S3 error: %s", err)
    }
}
