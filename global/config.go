package global


type ConfigDaemon struct {
	ListenSpec string
}
var CfgDaemon ConfigDaemon

type ConfigAuth struct {
	Login string
	Password string
	ShaPass string
}
var Auth ConfigAuth
