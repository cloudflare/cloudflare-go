package cloudflare

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAccessUserSeatWithMissingUID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessUserSeat(context.Background(), testAccountRC, UpdateAccessUserSeatParams{})
	assert.EqualError(t, err, "missing required access seat UID")

	_, err = client.UpdateAccessUserSeat(context.Background(), testZoneRC, UpdateAccessUserSeatParams{})
	assert.EqualError(t, err, "missing required access seat UID")
}
