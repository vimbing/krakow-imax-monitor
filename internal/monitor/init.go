package monitor

import (
	"github.com/vimbing/cclient"
	tls "github.com/vimbing/utls"
)

func Init() (*Monitor, error) {
	client, err := cclient.NewClient(tls.HelloChrome_100, "", true, 0)

	if err != nil {
		return &Monitor{}, err
	}

	return &Monitor{
		Client: &client,
	}, nil
}
