package HelloWorld1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityLog = logger.GetLogger("activity-flogo-HelloWorld1")

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
	logger.Debugf("Configuration Loaded")
	activityLog.Info("Configuration Loaded")
}

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	name := context.GetInput("name").(string)
	salutation := context.GetInput("salutation").(string)

	// Use the log object to log the greeting
	logger.Debugf("The Flogo engine says [%s] to [%s]", salutation, name)

	// Set the result as part of the context
	context.SetOutput("result", "The Flogo engine says "+salutation+" to "+name)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
