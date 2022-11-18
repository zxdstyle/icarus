package consoles

type Console interface {
	// Signature The name and signature of the consoles command.
	Signature() string
	// Description The consoles command description.
	Description() string
	// Handle Execute the consoles command.
	Handle(args ...string) error
}
