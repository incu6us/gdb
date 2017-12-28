gdb (GOLang Data Backup)
------

Tool for backup a local files, directories & mysql, using externally "mysqldump", "cp", "tar"

Config example:

```json
{
  "s3": {
    "bucket-endpoint": "https://test-bucket-01.nyc3.digitaloceanspaces.com",
    "region": "nyc3",
    "bucket": "test-bucket-01",
    "access-key": "*************",
    "secret-key": "**********************"
  },
  "mysql-configs": [
    {
      "host": "127.0.0.1",
      "port": 3306,
      "db": "database01",
      "user": "root",
      "password": "p@ssw0rd",
      "options": [],
      "s3-dir": "mysql-bkp"
    }
  ],
  "local-configs": [
    {
      "path": "/etc/issue",
      "s3-dir": "local-bkp"
    }
  ]
}
```

 * `s3` - general configuration for S3(for AWS S3 you could skip `bucket-endpoint` parameter in json config) or DigitalOcean Space;
 * `mysql-configs` - multiple configurations for MySQL database:
    -  `s3-dir` - path in S3 bucket where backup will be stored
 * `local-configs` - configurations for local files & directories to store it on S3