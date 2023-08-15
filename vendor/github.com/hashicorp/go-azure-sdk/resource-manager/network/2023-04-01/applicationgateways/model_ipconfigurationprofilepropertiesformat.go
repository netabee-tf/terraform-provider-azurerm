package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPConfigurationProfilePropertiesFormat struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Subnet            *Subnet            `json:"subnet,omitempty"`
}
