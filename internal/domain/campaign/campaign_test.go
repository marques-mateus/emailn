package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Content X"
	contacts := []string{"emai1@email.com", "email2@email.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal("1", campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}
