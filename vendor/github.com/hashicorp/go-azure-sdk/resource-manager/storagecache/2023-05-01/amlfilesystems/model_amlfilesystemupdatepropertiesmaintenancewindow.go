package amlfilesystems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmlFilesystemUpdatePropertiesMaintenanceWindow struct {
	DayOfWeek    *MaintenanceDayOfWeekType `json:"dayOfWeek,omitempty"`
	TimeOfDayUTC *string                   `json:"timeOfDayUTC,omitempty"`
}
