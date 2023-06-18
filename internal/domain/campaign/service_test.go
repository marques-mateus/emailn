package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type respositoryMock struct {
	mock.Mock
}

func (r *respositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Campaign X",
		Content: "Content X",
		Emails:  []string{"teste1@email.com", "teste2@email.com"},
	}
	service = Service{}
)

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign := newCampaign
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {

	repositoryMock := new(respositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name &&
			campaign.Content == newCampaign.Content &&
			len(campaign.Contacts) == len(newCampaign.Emails)
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(respositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
