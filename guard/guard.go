package guard

type (
	Authenticate interface {
		GetAuthIdentifier() uint
	}

	Guard interface {
		// Check Determine if the current user is authenticated.
		Check() error
		// ID Get the ID for the currently authenticated user.
		ID() uint
		// Validate a user's credentials.
		Validate() error
		// Login Log a user into the app.
		Login() any
	}
)
