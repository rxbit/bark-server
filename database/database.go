package database

// Database defines all of the db operation
type Database interface {
	CountAll() (int, error)                                 //Get db records count
	DeviceTokenByKey(key string) (string, error)            //Get specified device's token
	SaveDeviceTokenByKey(key, token string) (string, error) //Create or update specified devices's token
	Close() error                                           //Close the database
	SaveGroupByKeys(group_key string, keys []string) (string, error) //Create or update specified group's key
	GetDevicesByGroupKey(group_key string) ([]string, error) //Get specified group's devices
}
