package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

func TestGetPendingJobs(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/get_pending")
	// Act
	result, err := client.GetPendingJob("1234567890")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "public.ecr.aws/opslevel/cli:v2022.02.25", result.Image)
	autopilot.Equals(t, "ls -al", result.Commands[1])
}

func TestSetJobStatus(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/report_outcome")
	// Act
	err := client.ReportJobOutcome(ol.RunnerReportJobOutcomeInput{
		RunnerId:    ol.NewID("1234567890"),
		RunnerJobId: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvUnVubmVyczo6Sm9iUnVuLzE"),
		Outcome:     ol.RunnerJobOutcomeEnumExecutionTimeout,
	})
	// Assert
	autopilot.Ok(t, err)
}
