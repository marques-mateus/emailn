package database

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	r.campaigns = append(r.campaigns, *campaign)
	return nil
}

func (r *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return r.campaigns, nil
}

func (r *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	return nil, nil
}
