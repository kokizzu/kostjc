package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)

// Code generated by 1_codegen_test.go DO NOT EDIT.

var viewList = map[string]string{
	`Booking`:           `../svelte/booking.html`,           // ../svelte/booking.svelte
	`Building`:          `../svelte/building.html`,          // ../svelte/building.svelte
	`Debug`:             `../svelte/debug.html`,             // ../svelte/debug.svelte
	`Error`:             `../svelte/error.html`,             // ../svelte/error.svelte
	`Facility`:          `../svelte/facility.html`,          // ../svelte/facility.svelte
	`Index`:             `../svelte/index.html`,             // ../svelte/index.svelte
	`Location`:          `../svelte/location.html`,          // ../svelte/location.svelte
	`TenantsManagement`: `../svelte/tenantsManagement.html`, // ../svelte/tenantsManagement.svelte
}

func (v *Views) RenderBooking(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Booking`].Str(m))
}

func (v *Views) RenderBuilding(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Building`].Str(m))
}

func (v *Views) RenderDebug(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Debug`].Str(m))
}

func (v *Views) RenderError(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Error`].Str(m))
}

func (v *Views) RenderFacility(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Facility`].Str(m))
}

func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Index`].Str(m))
}

func (v *Views) RenderLocation(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Location`].Str(m))
}

func (v *Views) RenderTenantsManagement(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantsManagement`].Str(m))
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
