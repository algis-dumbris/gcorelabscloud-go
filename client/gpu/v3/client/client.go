package client

import (
	"github.com/urfave/cli/v2"

	gcorecloud "github.com/G-Core/gcorelabscloud-go"
	"github.com/G-Core/gcorelabscloud-go/client/common"
)

// NewGPUImageClientV3 creates a new GPU image client
func NewGPUImageClientV3(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	client, err := common.BuildClient(c, "gpu/baremetal", "v3")
	if err != nil {
		return nil, err
	}

	return client, nil
}
