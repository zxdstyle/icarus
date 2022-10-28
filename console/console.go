package console

type Console interface {
	// Signature The name and signature of the console command.
	Signature() string
	// Description The console command description.
	Description() string
	// Handle Execute the console command.
	Handle() error
}
