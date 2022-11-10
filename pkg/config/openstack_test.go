//go:build unit

package config_test

import (
	"testing"

	"github.com/squarefactory/cloud-burster/logger"
	"github.com/squarefactory/cloud-burster/pkg/config"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type OpenstackTestSuite struct {
	suite.Suite
}

func (suite *OpenstackTestSuite) TestValidate() {
	tests := []struct {
		input    string
		expected *config.Openstack
		isError  bool
		title    string
	}{
		{
			input: `enabled: true
identityEndpoint: 'https://auth.cloud.ovh.net/'
username: user
password: ''
region: GRA9
tenantID: tenantID
tenantName: 'tenantName'
domainID: default`,
			isError: false,
			expected: &config.Openstack{
				Enabled:          true,
				IdentityEndpoint: "https://auth.cloud.ovh.net/",
				UserName:         "user",
				Password:         "",
				TenantID:         "tenantID",
				TenantName:       "tenantName",
				DomainID:         "default",
				Region:           "GRA9",
			},
			title: "Positive test",
		},
		{
			input: `enabled: false
identityEndpoint: ''
username: ''
password: ''
region: ''
tenantID: ''
tenantName: ''
domainID: ''`,
			isError: false,
			expected: &config.Openstack{
				Enabled:          false,
				IdentityEndpoint: "",
				UserName:         "",
				Password:         "",
				TenantID:         "",
				TenantName:       "",
				DomainID:         "",
				Region:           "",
			},
			title: "Positive test",
		},
		{
			input: `enabled: true
identityEndpoint: 'aaa'
username: user
password: ''
region: GRA9
tenantID: tenantID
tenantName: 'tenantName'
domainID: default`,
			isError: true,
			expected: &config.Openstack{
				Enabled:          true,
				IdentityEndpoint: "aaa",
				UserName:         "user",
				Password:         "",
				TenantID:         "tenantID",
				TenantName:       "tenantName",
				DomainID:         "default",
				Region:           "GRA9",
			},
			title: "Valid URL",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.title, func() {
			// Arrange
			config := &config.Openstack{}
			err := yaml.Unmarshal([]byte(tt.input), config)
			suite.NoError(err)

			// Act
			err = config.Validate()

			// Assert
			if tt.isError {
				logger.I.Debug("expected error", zap.Error(err))
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
			suite.Equal(tt.expected, config)
		})
	}
}

func TestOpenstackTestSuite(t *testing.T) {
	suite.Run(t, &OpenstackTestSuite{})
}
