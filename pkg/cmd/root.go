package cmd

import (
	corecmd "github.com/Axway/agent-sdk/pkg/cmd"
	corecfg "github.com/Axway/agent-sdk/pkg/config"

	// CHANGE_HERE - Change the import path(s) below to reference packages correctly
	"github.com/simhead/amplify-3scale/pkg/config"
	"github.com/simhead/amplify-3scale/pkg/gateway"
)

// RootCmd - Agent root command
var RootCmd corecmd.AgentRootCmd
var gatewayConfig *config.GatewayConfig

func init() {
	// Create new root command with callbacks to initialize the agent config and command execution.
	// The first parameter identifies the name of the yaml file that agent will look for to load the config
	RootCmd = corecmd.NewRootCmd(
		"apic_discovery_agent",   // Name of the yaml file
		"3scale Discovery Agent", // Agent description
		initConfig,               // Callback for initializing the agent config
		run,                      // Callback for executing the agent
		corecfg.DiscoveryAgent,   // Agent Type (Discovery or Traceability)
	)

	// Get the root command properties and bind the config property in YAML definition
	rootProps := RootCmd.GetProperties()
	rootProps.AddStringProperty("3scale-api-gateway.specPath", "./apis/musical_instruments.json", "Sample Swagger specification path for discovery")
	rootProps.AddStringProperty("3scale-api-gateway.config_key_1", "", "Config Key 1")
	rootProps.AddStringProperty("3scale-api-gateway.config_key_2", "", "Config Key 1")
	rootProps.AddStringProperty("3scale-api-gateway.config_key_3", "", "Config Key 3")

}

// Callback that agent will call to process the execution
func run() error {
	gatewayClient, err := gateway.NewClient(gatewayConfig)
	err = gatewayClient.DiscoverAPIs()
	return err
}

// Callback that agent will call to initialize the config. CentralConfig is parsed by Agent SDK
// and passed to the callback allowing the agent code to access the central config
func initConfig(centralConfig corecfg.CentralConfig) (interface{}, error) {
	rootProps := RootCmd.GetProperties()
	// Parse the config from bound properties and setup gateway config
	gatewayConfig = &config.GatewayConfig{
		SpecPath:   rootProps.StringPropertyValue("3scale-api-gateway.specPath"),
		ConfigKey1: rootProps.StringPropertyValue("3scale-api-gateway.config_key_1"),
		ConfigKey2: rootProps.StringPropertyValue("3scale-api-gateway.config_key_2"),
		ConfigKey3: rootProps.StringPropertyValue("3scale-api-gateway.config_key_3"),
	}

	agentConfig := config.AgentConfig{
		CentralCfg: centralConfig,
		GatewayCfg: gatewayConfig,
	}
	return agentConfig, nil
}

// GetAgentConfig - Returns the agent config
func GetAgentConfig() *config.GatewayConfig {
	return gatewayConfig
}
