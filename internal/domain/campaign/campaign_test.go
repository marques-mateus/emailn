package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Content X"
	contacts = []string{"emai1@email.com", "email2@email.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_NameIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.Name)
}

func Test_NewCampaign_ContentIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.Content)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
