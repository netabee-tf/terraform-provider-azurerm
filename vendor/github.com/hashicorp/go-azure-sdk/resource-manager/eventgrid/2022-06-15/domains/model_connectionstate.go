package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionState struct {
	ActionsRequired *string                    `json:"actionsRequired,omitempty"`
	Description     *string                    `json:"description,omitempty"`
	Status          *PersistedConnectionStatus `json:"status,omitempty"`
}
