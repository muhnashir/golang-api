package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct{
	ID	int
	UserID int
	Name string
	ShortDescription string
	Description string
	Perks string
	BackerCount int
	GoalAmount int
	CurrentAmount int
	Slug string
	CreatedAt time.Time
	UpdatedAt time.Time
	CampaignImages []CampaignImage
	User user.User
}

type CampaignImage struct{
	ID	int
	CampaignID int
	FileName string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCampaignInput struct{
	Name string	`json:"name"`
	ShortDescription string	`json:"short_description"`
	Description string	`json:"description"`
	Perks string	`json:"perks"`
	GoalAmount int	`json:"goal_amount"`
	Slug string	`json:"slug"`
	User user.User	`json:"user"`
}