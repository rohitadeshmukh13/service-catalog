// Copyright 2014-2016 Fraunhofer Institute for Applied Information Technology FIT

package catalog

const (
	DNSSDServiceType = "_linksmart-sc._tcp"
	MaxPerPage       = 100
	APIVersion       = "2.0.0"
	LoggerPrefix     = "[sc] "

	CatalogBackendMemory  = "memory"
	CatalogBackendLevelDB = "leveldb"
)

var SupportedBackends = map[string]bool{
	CatalogBackendMemory:  true,
	CatalogBackendLevelDB: true,
}

var SupportedProtocols = map[string]bool{
	"HTTP": true,
	"MQTT": true,
	"AMQP": true,
}
