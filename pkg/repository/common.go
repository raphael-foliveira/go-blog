package repository

import (
	"fmt"
	"strings"
)

func buildQuery(query string, conditions ...string) string {
	if len(conditions) > 0 {
		query += ` WHERE `
		for _, c := range conditions {
			query += fmt.Sprintf(" %v ", strings.TrimSpace(c))
		}
		query += `;`
	}
	return query
}
