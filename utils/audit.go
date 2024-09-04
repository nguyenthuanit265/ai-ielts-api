package utils

import (
	"context"
	"time"
)

func AuditUpdateData(ctx context.Context) (int, int, time.Time, time.Time) {
	currentUser, _ := GetCurrentUser(ctx)
	// Build request insert
	createdBy := currentUser.Id
	updatedBy := currentUser.Id
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()

	return createdBy, updatedBy, createdAt, updatedAt
}
