package visionsummary

import "testing"

func TestGenerateSummaryFromKeywords(t *testing.T) {
	s := GenerateSummaryFromKeywords([]string{"Audio", "Space", "Physics"})
	if len(s.KeyTopics) != 3 {
		t.Errorf("unexpected topics count: %+v", s)
	}
}
