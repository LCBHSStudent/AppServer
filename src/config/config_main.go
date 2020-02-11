package uconfig

import "time"

const ProjectPath string = "A:/WorkSpace/GoProject/AppServer/"

const KeyPath string = ProjectPath + "key_config/default.key"
const PemPath string = ProjectPath + "key_config/default.pem"

const ErrCertInvalid int = 0x10000000

const ServiceReadTimeout  time.Duration = time.Second * 15
const ServiceWriteTimeout time.Duration = time.Second * 10

const GatewayIdleTimeout time.Duration = time.Second * 10
const GatewayMaxConnection 	       int = 100


const ServicePort int = 1919
