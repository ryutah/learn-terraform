package main

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "apply_helloworld",
			expected: "Hello, World!!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			terraformOption := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "./",
			})
			defer terraform.Destroy(t, terraformOption)

			terraform.InitAndApply(t, terraformOption)

			output := terraform.Output(t, terraformOption, "hello_world")
			assert.Equal(t, test.expected, output)
		})
	}
}
