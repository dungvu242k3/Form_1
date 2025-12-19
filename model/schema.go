package models

import "time"

type Address struct {
	City     string `bson:"city,omitempty"`
	District string `bson:"district,omitempty"`
	Street   string `bson:"street,omitempty"`
}

type PersonalInfo struct {
	FullName    string  `bson:"fullName"`
	DateOfBirth string  `bson:"dateOfBirth,omitempty"`
	Age         int     `bson:"age,omitempty"`
	Nationality string  `bson:"nationality,omitempty"`
	Email       string  `bson:"email,omitempty"`
	Phone       string  `bson:"phone,omitempty"`
	Address     Address `bson:"address,omitempty"`
}

type PhysicalInfo struct {
	Height    float64 `bson:"height,omitempty"`
	Weight    float64 `bson:"weight,omitempty"`
	Bust      float64 `bson:"bust,omitempty"`
	Waist     float64 `bson:"waist,omitempty"`
	Hips      float64 `bson:"hips,omitempty"`
	EyeColor  string  `bson:"eyeColor,omitempty"`
	HairColor string  `bson:"hairColor,omitempty"`
}

type SkillsAndEducation struct {
	Education  string   `bson:"education,omitempty"`
	Languages  []string `bson:"languages,omitempty"`
	Skills     []string `bson:"skills,omitempty"`
	Experience []string `bson:"experience,omitempty"`
}

type SocialLinks struct {
	Instagram string `bson:"instagram,omitempty"`
	Facebook  string `bson:"facebook,omitempty"`
	TikTok    string `bson:"tiktok,omitempty"`
}

type Portfolio struct {
	ProfilePicture string      `bson:"profilePicture,omitempty"`
	PortfolioLinks []string    `bson:"portfolioLinks,omitempty"`
	Biography      string      `bson:"biography,omitempty"`
	SocialLinks    SocialLinks `bson:"socialLinks,omitempty"`
}

type Management struct {
	IsActive         bool      `bson:"isActive"`
	IsVisibleToUsers bool      `bson:"isVisibleToUsers,omitempty"`
	SubmittedAt      time.Time `bson:"submittedAt,omitempty"`
	UpdatedAt        time.Time `bson:"updatedAt,omitempty"`
}

type Contestant struct {
	ID                 string             `bson:"_id,omitempty"`
	PersonalInfo       PersonalInfo       `bson:"personalInfo"`
	PhysicalInfo       PhysicalInfo       `bson:"physicalInfo"`
	SkillsAndEducation SkillsAndEducation `bson:"skillsAndEducation"`
	Portfolio          Portfolio          `bson:"portfolio"`
	Management         Management         `bson:"management"`
}
