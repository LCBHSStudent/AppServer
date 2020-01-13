package uconfig

import "time"

const KeyPath string = "A:/WorkSpace/GoProject/AppServer/key_config/default.key"
const PemPath string = "A:/WorkSpace/GoProject/AppServer/key_config/default.pem"

const ErrCertInvalid int = 0x10000000

const ServiceReadTimeout  time.Duration = time.Second * 15
const ServiceWriteTimeout time.Duration = time.Second * 10

const ServicePort int = 1919