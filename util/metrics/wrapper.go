package metrics

import (
	"github.com/Rahul12344/skelego"
)

//LogWithSteps Runs a workflow with log emissions/step
func LogWithSteps(logger skelego.Logging, workflowName string, stepName string, workflow func() (interface{}, error)) interface{} {
	logger.LogEvent("Running workflow: %s, step: %s", workflowName, stepName)
	output, err := workflow()
	if err != nil {
		logger.LogError("Workflow: %s terminated with error: ", err.Error())
	}
	logger.LogEvent("Finished workflow: %s", workflowName)
	return output
}
