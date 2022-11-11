//go:build unit

package config_test

import (
	"testing"

	"github.com/squarefactory/cloud-burster/pkg/config"
	"github.com/stretchr/testify/suite"
)

var cleanOpenstackCloud = config.Cloud{
	AuthorizedKeys: []string{
		"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDUnXMBGq6bV6H+c7P5QjDn1soeB6vkodi6OswcZsMwH nguye@PC-DARKNESS4",
	},
	PostScripts: config.PostScriptsOpts{
		Git: config.GitOpts{
			Key: "key",
			URL: "git@github.com:SquareFactory/compute-configs.git",
			Ref: "main",
		},
	},
	Network: cleanNetwork,
	Hosts: []config.Host{
		cleanHost,
	},
	GroupsHost: []config.GroupHost{
		cleanGroupHost,
	},
	Type:      "openstack",
	Openstack: &cleanOpenstack,
}

var cleanExoscaleCloud = config.Cloud{
	AuthorizedKeys: []string{
		"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDUnXMBGq6bV6H+c7P5QjDn1soeB6vkodi6OswcZsMwH nguye@PC-DARKNESS4",
	},
	PostScripts: config.PostScriptsOpts{
		Git: config.GitOpts{
			Key: "key",
			URL: "git@github.com:SquareFactory/compute-configs.git",
			Ref: "main",
		},
	},
	Network: cleanNetwork,
	Hosts: []config.Host{
		cleanHost,
	},
	GroupsHost: []config.GroupHost{
		cleanGroupHost,
	},
	Type:     "exoscale",
	Exoscale: &cleanExoscale,
}

type CloudTestSuite struct {
	suite.Suite
}

func (suite *CloudTestSuite) TestValidate() {
	tests := []struct {
		input         *config.Cloud
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: &cleanOpenstackCloud,
			title: "Positive test",
		},
		{
			isError: true,
			errorContains: []string{
				"required",
				"Network",
			},
			input: &config.Cloud{
				AuthorizedKeys: cleanOpenstackCloud.AuthorizedKeys,
				PostScripts:    cleanOpenstackCloud.PostScripts,
				GroupsHost:     cleanOpenstackCloud.GroupsHost,
				Openstack:      cleanOpenstackCloud.Openstack,
				Type:           cleanOpenstackCloud.Type,
			},
			title: "Network required/valid",
		},
		{
			isError: true,
			errorContains: []string{
				"GroupsHost",
			},
			input: &config.Cloud{
				AuthorizedKeys: cleanOpenstackCloud.AuthorizedKeys,
				PostScripts:    cleanOpenstackCloud.PostScripts,
				Network:        cleanOpenstackCloud.Network,
				GroupsHost: []config.GroupHost{
					{},
				},
				Openstack: cleanOpenstackCloud.Openstack,
				Type:      cleanOpenstackCloud.Type,
			},
			title: "GroupsHost valid",
		},
		{
			isError: true,
			errorContains: []string{
				"Openstack",
			},
			input: &config.Cloud{
				AuthorizedKeys: cleanOpenstackCloud.AuthorizedKeys,
				PostScripts:    cleanOpenstackCloud.PostScripts,
				Network:        cleanOpenstackCloud.Network,
				GroupsHost:     cleanOpenstackCloud.GroupsHost,
				Type:           cleanOpenstackCloud.Type,
				Openstack: &config.Openstack{
					IdentityEndpoint: "aaa",
				},
			},
			title: "openstack valid",
		},
		{
			isError: true,
			errorContains: []string{
				"required_if",
				"Openstack",
			},
			input: &config.Cloud{
				AuthorizedKeys: cleanOpenstackCloud.AuthorizedKeys,
				PostScripts:    cleanOpenstackCloud.PostScripts,
				Network:        cleanOpenstackCloud.Network,
				GroupsHost:     cleanOpenstackCloud.GroupsHost,
				Type:           "openstack",
			},
			title: "If type == openstack, openstack is required",
		},
		{
			isError: true,
			errorContains: []string{
				"required_if",
				"Exoscale",
			},
			input: &config.Cloud{
				AuthorizedKeys: cleanOpenstackCloud.AuthorizedKeys,
				PostScripts:    cleanOpenstackCloud.PostScripts,
				Network:        cleanOpenstackCloud.Network,
				GroupsHost:     cleanOpenstackCloud.GroupsHost,
				Type:           "exoscale",
			},
			title: "If type == exoscale, exoscale is required",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.title, func() {
			// Act
			err := tt.input.Validate()

			// Assert
			if tt.isError {
				suite.Error(err)
				for _, contain := range tt.errorContains {
					suite.ErrorContains(err, contain)
				}
			} else {
				suite.NoError(err)
			}
		})
	}
}

func TestCloudTestSuite(t *testing.T) {
	suite.Run(t, &CloudTestSuite{})
}
