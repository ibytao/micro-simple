package jwt

import "micro-simple/basic"

type Jwt struct {
	SecretKey string `json:"secretKey"`
}

func init() {
	basic.Register(initJwt)
}

func initJwt() {}
