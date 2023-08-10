package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	OSS struct {
		AccessKey  string
		SecretKey  string
		BucketName string
		DomainName string
	}
}
