// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package contentservice

import (
	"github.com/gitpod-io/gitpod/usage/pkg/db"
	"time"
)

type InvalidSession struct {
	Reason  string
	Session db.WorkspaceInstanceForUsage
}

type UsageReport struct {
	GenerationTime time.Time `json:"generation_time"`

	From time.Time `json:"from"`
	To   time.Time `json:"to"`

	RawSessions     []db.WorkspaceInstanceForUsage `json:"rawSessions"`
	InvalidSessions []InvalidSession               `json:"invalidSessions"`

	UsageRecords []db.WorkspaceInstanceUsage `json:"usageRecords"`
}
