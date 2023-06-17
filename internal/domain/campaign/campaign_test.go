package campaign

import (
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Content X"
	contacts := []string{"emai1@email.com", "email2@email.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("campaign.ID = %v; want %v", campaign.ID, "1")
	}

	if campaign.Name != name {
		t.Errorf("campaign.Name = %v; want %v", campaign.Name, name)
	}

	if campaign.Content != content {
		t.Errorf("campaign.Content = %v; want %v", campaign.Content, content)
	}

	if len(campaign.Contacts) != len(contacts) {
		t.Errorf("campaign.Contacts = %v; want %v", len(campaign.Contacts), len(contacts))
	}
}
