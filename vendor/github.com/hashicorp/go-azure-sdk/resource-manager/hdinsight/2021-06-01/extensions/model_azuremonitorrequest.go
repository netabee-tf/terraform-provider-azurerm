package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureMonitorRequest struct {
	PrimaryKey             *string                             `json:"primaryKey,omitempty"`
	SelectedConfigurations *AzureMonitorSelectedConfigurations `json:"selectedConfigurations,omitempty"`
	WorkspaceId            *string                             `json:"workspaceId,omitempty"`
}
