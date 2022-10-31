package guard

import "github.com/zxdstyle/icarus/server/requests"

type (
	Authenticate interface {
		GetAuthIdentifier() uint
	}

	Guard interface {
		// Check Determine if the current user is authenticated.
		Check(req requests.Request) error
		// ID Get the ID for the currently authenticated user.
		ID(req requests.Request) uint
		// LoginUsingID Log a user into the app.
		LoginUsingID(id uint) (any, error)
	}
)
