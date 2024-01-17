package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
)

const (
	ConfigPath = "cloud"

	AzurePublicName       = "AzureCloud"
	AzureChinaCloudName   = "AzureChinaCloud"
	AzureUSGovernmentName = "AzureUSGovernment"
)

type Cloud struct {
	Configuration cloud.Configuration

	// The base URL for the cloud's portal (e.g. https://portal.azure.com for
	// Azure public cloud).
	PortalUrlBase string
}

type Config struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

func NewCloud(config *Config) *Cloud {
	if cloud, err := getNamedCloud(config.Name); err != nil {
		// TODO: Find a friendly way to surface errors here
		panic(err)
	} else {
		return cloud
	}
}

func ParseCloudConfig(partialConfig any) (*Config, error) {
	var config *Config

	jsonBytes, err := json.Marshal(partialConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cloud configuration: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cloud configuration: %w", err)
	}

	return config, nil
}

func GetAzurePublic() *Cloud {
	return &Cloud{
		Configuration: cloud.AzurePublic,
		PortalUrlBase: "https://portal.azure.com",
	}
}

func GetAzureGovernment() *Cloud {
	return &Cloud{
		Configuration: cloud.AzureGovernment,
		PortalUrlBase: "https://portal.azure.us",
	}
}

func GetAzureChina() *Cloud {
	return &Cloud{
		Configuration: cloud.AzureChina,
		PortalUrlBase: "https://portal.azure.cn",
	}
}

func getNamedCloud(name string) (*Cloud, error) {
	if name == AzurePublicName || name == "" {
		return GetAzurePublic(), nil
	} else if name == AzureChinaCloudName {
		return GetAzureChina(), nil
	} else if name == AzureUSGovernmentName {
		return GetAzureGovernment(), nil
	}

	return &Cloud{}, fmt.Errorf(
		"cloud name '%s' not found use valid names '%s', '%s', '%s'",
		name,
		AzurePublicName,
		AzureChinaCloudName,
		AzureUSGovernmentName,
	)
}