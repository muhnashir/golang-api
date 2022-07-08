package campaign

type CampaignFormatter struct{
	ID int `json:"id"`
	UserID int	`json:"user_id"`
	Name string	`json:"name"`
	ShortDescription string	`json:"short_description"`
	ImageURL string	`json:"image_url"`
	GoalAmount int	`json:"goal_amount"`
	CurrentAmout int	`json:"current_amount"`
	Slug string `json:"slug"`
}

func FormatCampaign(campaign Campaign) (CampaignFormatter){
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmout = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if(len(campaign.CampaignImages) > 0){
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	
	return campaignFormatter
}


func FormatCampaigns(campaigns []Campaign)[]CampaignFormatter{
	sliceCampaignFormatter := []CampaignFormatter{}


	for _, campaign := range campaigns{
		campaignFormatter := FormatCampaign(campaign)
		sliceCampaignFormatter = append(sliceCampaignFormatter, campaignFormatter)
	}

	return sliceCampaignFormatter
}

