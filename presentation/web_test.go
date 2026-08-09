package presentation

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"kostjc/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalWebJSONAcceptsNumbersForJSONStringTaggedFields(t *testing.T) {
	var in domain.StaffPaymentPromiseIn

	err := unmarshalWebJSON([]byte(`{
		"cmd": "upsert",
		"promise": {
			"tenantId": 123,
			"bookingId": 456,
			"promisedAt": "2026-07-25",
			"promisePayAt": "2026-07-29",
			"note": "payday"
		}
	}`), &in)

	require.NoError(t, err)
	require.Equal(t, uint64(123), in.Promise.TenantId)
	require.Equal(t, uint64(456), in.Promise.BookingId)
	require.Equal(t, `2026-07-25`, in.Promise.PromisedAt)
	require.Equal(t, `2026-07-29`, in.Promise.PromisePayAt)
	require.Equal(t, `payday`, in.Promise.Note)
}

func TestUnmarshalWebJSONNormalizesJSONStringTaggedFieldsInSlices(t *testing.T) {
	type item struct {
		TenantId uint64 `json:"tenantId,string"`
		Count    uint64 `json:"count"`
	}
	type input struct {
		Items []item `json:"items"`
	}

	var in input
	err := unmarshalWebJSON([]byte(`{
		"items": [
			{ "tenantId": 123, "count": 456 },
			{ "tenantId": "789", "count": 101 }
		]
	}`), &in)

	require.NoError(t, err)
	require.Equal(t, []item{
		{TenantId: 123, Count: 456},
		{TenantId: 789, Count: 101},
	}, in.Items)
}

func TestWebApiParseInputAcceptsNumericJSONStringTaggedFields(t *testing.T) {
	app := fiber.New()
	app.Post(`/promise`, func(c *fiber.Ctx) error {
		var in domain.StaffPaymentPromiseIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.StaffPaymentPromiseAction)
		require.NoError(t, err)
		require.Equal(t, uint64(123), in.Promise.TenantId)
		require.Equal(t, uint64(456), in.Promise.BookingId)
		return c.SendStatus(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, `/promise`, strings.NewReader(`{
		"cmd": "upsert",
		"promise": {
			"tenantId": 123,
			"bookingId": 456,
			"promisedAt": "2026-07-25",
			"promisePayAt": "2026-07-29"
		}
	}`))
	req.Header.Set(fiber.HeaderContentType, `application/json; charset=utf-8`)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)

	req = httptest.NewRequest(http.MethodPost, `/promise`, strings.NewReader(`{
		"cmd": "upsert",
		"promise": {
			"TenantId": 123,
			"BookingId": 456,
			"promisedAt": "2026-07-25",
			"promisePayAt": "2026-07-29"
		}
	}`))
	req.Header.Set(fiber.HeaderContentType, `application/json`)

	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}
