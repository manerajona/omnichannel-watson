# Watson Connector

## api.go 
Implements the Watson API. We use the interfaces provided by the official Watson SDK for Go (https://github.com/watson-developer-cloud/go-sdk). The SDK allows us to create and delete sessions, as well as send an input to our Watson Assistant instance.

## config.go 
Allows us to specify a path to read a .yaml file, in this file we store the configuration data of our assistant such as: the assistant id, the instance id, the region, the api-key and the version.