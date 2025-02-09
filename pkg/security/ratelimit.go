package security

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type RateLimiterConfig struct {
	Max        int
	Expiration time.Duration
}

func SetupRateLimiter(app *fiber.App, rateLimiterConfig RateLimiterConfig) {

	if rateLimiterConfig.Max == 0 {
		rateLimiterConfig.Max = 20
	}
	if rateLimiterConfig.Expiration == 0 {
		rateLimiterConfig.Expiration = 30 * time.Second
	}

	app.Use(limiter.New(limiter.Config{
		Max:        rateLimiterConfig.Max,
		Expiration: rateLimiterConfig.Expiration,
	}))
}
