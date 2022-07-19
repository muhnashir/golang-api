package campaign

import "bwastartup/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string    `json:"name" binding:"required"`
	ShortDescription string    `json:"short_description" binding:"required"`
	Description      string    `json:"description" binding:"required"`
	Perks            string    `json:"perks" binding:"required"`
	GoalAmount       int       `json:"goal_amount" binding:"required"`
	Slug             string    `json:"slug"`
	User             user.User `json:"user"`
}

type CreateCampaignImageInput struct {
	CampaignID int `form:"campaign_id" binding:"required"`
	isPrimary  bool   `form:"is_primary" binding:"required"`
}
