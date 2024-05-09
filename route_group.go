package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritd/logger"
)

type GroupInfo struct {
	GroupKey   string `form:"group_key,omitempty" json:"group_key,omitempty" xml:"group_key,omitempty" query:"group_key,omitempty"`
	DeviceKeys   []string `form:"device_keys,omitempty" json:"device_keys,omitempty" xml:"device_keys,omitempty" query:"device_keys,omitempty"`
}

func init() {
	registerRoute("group", func(router fiber.Router) {
		router.Post("/group", func(c *fiber.Ctx) error { return doRegisterGroup(c, false) })
	})
}

func doRegisterGroup(c *fiber.Ctx, compat bool) error {
	var deviceInfo GroupInfo
	if err := c.BodyParser(&deviceInfo); err != nil {
		return c.Status(400).JSON(failed(400, "request group failed: %v", err))
	}

	if (deviceInfo.DeviceKeys == nil || len(deviceInfo.DeviceKeys) == 0) {
		return c.Status(400).JSON(failed(400, "request group failed: device_keys is empty"))
	}

	group_key, err := db.SaveGroupByKeys(deviceInfo.GroupKey, deviceInfo.DeviceKeys)
	if err != nil {
		logger.Errorf("device group failed: %v", err)
		return c.Status(500).JSON(failed(500, "device group failed: %v", err))
	}

	return c.Status(200).JSON(data(map[string]interface{}{
		// compatible with old resp
		"group_key": group_key,
		"device_keys": deviceInfo.DeviceKeys,
	}))
}
