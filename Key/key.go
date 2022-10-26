package Key

import (
	"AI/conf"
)

var Key conf.Settings

func InitKey(settings conf.Settings) {
	Key.ApiKey = settings.ApiKey
	Key.SecreatKey = settings.SecreatKey
}