package security

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiterConfig is the configuration for rate limiting
type RateLimiterConfig struct {
	Max        int           // Max requests allowed
	Expiration time.Duration // Expiration time
}

// SetupRateLimiter sets up the rate limiting middleware with custom configuration
func SetupRateLimiter(app *fiber.App, rateLimiterConfig RateLimiterConfig) {
	// Set default values if config is not set
	if rateLimiterConfig.Max == 0 {
		rateLimiterConfig.Max = 20 // Default max requests
	}
	if rateLimiterConfig.Expiration == 0 {
		rateLimiterConfig.Expiration = 30 * time.Second // Default expiration time
	}

	// Apply rate limiter middleware to the app
	app.Use(limiter.New(limiter.Config{
		Max:        rateLimiterConfig.Max,
		Expiration: rateLimiterConfig.Expiration,
	}))
}
