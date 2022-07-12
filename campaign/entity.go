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
	Name string	`json:"name" binding:"required"`
	ShortDescription string	`json:"short_description" binding:"required"`
	Description string	`json:"description" binding:"required"`
	Perks string	`json:"perks" binding:"required"`
	GoalAmount int	`json:"goal_amount" binding:"required"`
	Slug string	`json:"slug"`
	User user.User	`json:"user"`
}