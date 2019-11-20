package yalp

// Logger is a wrapper for any kind of logger you wish to use. This can be
// customized and changed within the Quirk client itself.
type Logger interface {
	Info(msg string, fields ...interface{})
	// Debug is used when confirming when things are doing their jobs such as
	// when adding vertex labels to the schema.
	Debug(msg string, fields ...interface{})
	// Warn is a lot like Error in that it does not stop the application, but
	// is logged at a different level to warn the developer or user of something
	// wrong or something to do later.
	Warn(msg string, fields ...interface{})
	// Error is used when there is a problem but not a big enough problem to
	// stop an app. These problems are minor, but not major.
	Error(msg string, fields ...interface{})
	// Fatal's purpose is to stop the application because something really wrong
	// happened. A case of this being used is when trying to put an odd number
	// of properties in an AddVertex function. Which would not create a proper
	// query for the gremlin server and should stop.
	Fatal(msg string, fields ...interface{})
	// Sugar returns a sugared logger, which is typically slower for logging
	// but is nicer to read than the typical logger.
	Sugar() Logger
}
