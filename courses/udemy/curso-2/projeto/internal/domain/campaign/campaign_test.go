package campaign_test

import (
	"email/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	//Proprieties
	name     = "Campaign-X"
	content  = "Body"
	contacts = []string{"email@gmail.com", "outroemail@gmail.com"}
)

func TestNewCampaign(t *testing.T) {
	//Assert
	assert := assert.New(t)

	//Define campaign
	campaign, _ := campaign.NewCampaign(name, content, contacts)

	//Testing value in campaign proprieties
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

	//Testing if not nil
	assert.NotNil(campaign.ID)

	//Testing if greater than time.now less one minute
	assert.Greater(campaign.CreatedAt, time.Now().Add(-time.Minute))

}

func TestErrorsCampaign(t *testing.T) {
	assert := assert.New(t)
	_, e := campaign.NewCampaign(name, content, contacts)

	assert.Empty(e)

}
