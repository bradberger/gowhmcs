package whmcs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetRegDate(t *testing.T) {
	var d ClientDomainUpdateReq
	tm := time.Date(2016, 01, 17, 9, 30, 0, 0, time.UTC)
	d.SetRegDate(tm)
	d.SetNextDueDate(tm)
	d.SetExpiryDate(tm)
	assert.Equal(t, d.RegDate, "2016-01-17")
	assert.Equal(t, d.ExpiryDate, "2016-01-17")
	assert.Equal(t, d.NextDueDate, "2016-01-17")
}
