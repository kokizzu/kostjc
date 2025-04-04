package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


var viewList = map[string]string{
	`Debug`: `../svelte/debug.html`, // ../svelte/debug.svelte
	`Error`: `../svelte/error.html`, // ../svelte/error.svelte
	`Index`: `../svelte/index.html`, // ../svelte/index.svelte
}


func (v *Views) RenderDebug(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Debug`].Str(m))
}

func (v *Views) RenderError(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Error`].Str(m))
}

func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Index`].Str(m))
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
