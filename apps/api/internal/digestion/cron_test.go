package digestion

import "testing"

func TestCronStartStop(t *testing.T) {
	cron := NewCron()
	cron.Start()
	if cron.ticker == nil {
		t.Fatalf("expected ticker to be initialized")
	}
	cron.Stop()
}

func TestCronStopWithoutStart(t *testing.T) {
	cron := NewCron()
	cron.Stop()
}
