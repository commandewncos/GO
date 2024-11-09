package campaign_test

import (
	"email/internal/contract"
	"email/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {

	args := r.Called(campaign)
	return args.Error(0)

}

func Test_Create_Capaign(t *testing.T) {
	// assert := assert.New(t)
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Content Y",
		Emails:  []string{"one@gmail.com", "two@gmail.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name {

			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service := campaign.Service{repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}
