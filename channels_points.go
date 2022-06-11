package helix

type ChannelCustomRewardsParams struct {
	BroadcasterID                     string `query:"broadcaster_id"`
	Title                             string `json:"title"`
	Cost                              int    `json:"cost"`
	Prompt                            string `json:"prompt"`
	IsEnabled                         bool   `json:"is_enabled"`
	BackgroundColor                   string `json:"background_color"`
	IsUserInputRequired               bool   `json:"is_user_input_required"`
	IsMaxPerStreamEnabled             bool   `json:"is_max_per_stream_enabled"`
	MaxPerStream                      int    `json:"max_per_stream"`
	IsMaxPerUserPerStreamEnabled      bool   `json:"is_max_per_user_per_stream_enabled"`
	MaxPerUserPerStream               int    `json:"max_per_user_per_stream"`
	IsGlobalCooldownEnabled           bool   `json:"is_global_cooldown_enabled"`
	GlobalCooldownSeconds             int    `json:"global_cooldown_seconds"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue"`
}

type DeleteCustomRewardsParams struct {
	BroadcasterID string `query:"broadcaster_id"`
	ID            string `query:"id"`
}

type UpdateCustomRewardsParams struct {
	BroadcasterID string `query:"broadcaster_id"`
	Id            string `query:"id"`
	Title         string `json:"title"`
	Cost          int    `json:"cost"`
	Prompt        string `json:"prompt"`
	IsEnabled     bool   `json:"is_enabled"`
}

type GetCustomRewardsParams struct {
	BroadcasterID         string `query:"broadcaster_id"`
	ID                    string `query:"id"`
	OnlyManageableRewards bool   `query:"only_manageable_rewards"`
}

type GetCustomRewardsRedemptionParams struct {
	BroadcasterID string `query:"broadcaster_id"`
	RewardId      string `query:"reward_id"`
	ID            string `query:"id"`
	Status        string `query:"status"`
	After         string `query:"after"`
	First         int    `query:"first"`
}

type ManyChannelCustomRewards struct {
	ChannelCustomRewards []ChannelCustomReward `json:"data"`
}

type ChannelCustomReward struct {
	BroadcasterID                     string                      `json:"broadcaster_id"`
	BroadcasterLogin                  string                      `json:"broadcaster_login"`
	BroadcasterName                   string                      `json:"broadcaster_name"`
	ID                                string                      `json:"id"`
	Title                             string                      `json:"title"`
	Prompt                            string                      `json:"prompt"`
	Cost                              int                         `json:"cost"`
	Image                             RewardImage                 `json:"image"`
	DefaultImage                      RewardImage                 `json:"default_image"`
	IsEnabled                         bool                        `json:"is_enabled"`
	IsUserInputRequired               bool                        `json:"is_user_input_required"`
	MaxPerStreamSetting               MaxPerStreamSettings        `json:"max_per_stream_setting"`
	MaxPerUserPerStreamSetting        MaxPerUserPerStreamSettings `json:"max_per_user_per_stream_setting"`
	GlobalCooldownSetting             GlobalCooldownSettings      `json:"global_cooldown_setting"`
	IsPaused                          bool                        `json:"is_paused"`
	IsInStock                         bool                        `json:"is_in_stock"`
	ShouldRedemptionsSkipRequestQueue bool                        `json:"should_redemptions_skip_request_queue"`
	RedemptionsRedeemedCurrentStream  int                         `json:"redemptions_redeemed_current_stream"`
	CooldownExpiresAt                 string                      `json:"cooldown_expires_at"`
}

type ManyCustomRewardRedemption struct {
	CustomRewardRedemption []CustomRewardRedemption `json:"data"`
	Pagination             *Pagination              `json:"pagination,omitempty"`
}

type CustomRewardRedemption struct {
	BroadcasterName  string  `json:"broadcaster_name"`
	BroadcasterLogin string  `json:"broadcaster_login"`
	BroadcasterID    string  `json:"broadcaster_id"`
	ID               string  `json:"id"`
	UserLogin        *string `json:"user_login,omitempty"`
	UserID           string  `json:"user_id"`
	UserName         string  `json:"user_name"`
	UserInput        string  `json:"user_input"`
	Status           string  `json:"status"`
	RedeemedAt       string  `json:"redeemed_at"`
	Reward           Reward  `json:"reward"`
}

type Reward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Prompt string `json:"prompt"`
	Cost   int64  `json:"cost"`
}

type RewardImage struct {
	Url1x string `json:"url_1x"`
	Url2x string `json:"url_2x"`
	Url4x string `json:"url_4x"`
}

type MaxPerUserPerStreamSettings struct {
	IsEnabled           bool `json:"is_enabled"`
	MaxPerUserPerStream int  `json:"max_per_user_per_stream"`
}

type MaxPerStreamSettings struct {
	IsEnabled    bool `json:"is_enabled"`
	MaxPerStream int  `json:"max_per_stream"`
}

type GlobalCooldownSettings struct {
	IsEnabled             bool `json:"is_enabled"`
	GlobalCooldownSeconds int  `json:"global_cooldown_seconds"`
}

type ChannelCustomRewardResponse struct {
	ResponseCommon
	Data ManyChannelCustomRewards
}

type RedemptionResponse struct {
	ResponseCommon
	Data ManyCustomRewardRedemption
}

// Response for removing a custom reward
type DeleteCustomRewardsResponse struct {
	ResponseCommon
}

// CreateCustomReward : Creates a Custom Reward on a channel.
// Required scope: channel:manage:redemptions
func (c *Client) CreateCustomReward(params *ChannelCustomRewardsParams) (*ChannelCustomRewardResponse, error) {
	resp, err := c.postAsJSON("/channel_points/custom_rewards", &ManyChannelCustomRewards{}, params)
	if err != nil {
		return nil, err
	}

	reward := &ChannelCustomRewardResponse{}
	resp.HydrateResponseCommon(&reward.ResponseCommon)
	reward.Data.ChannelCustomRewards = resp.Data.(*ManyChannelCustomRewards).ChannelCustomRewards

	return reward, nil
}

// DeleteCustomRewards : Deletes a Custom Rewards on a channel
// Required scope: channel:manage:redemptions
func (c *Client) DeleteCustomRewards(params *DeleteCustomRewardsParams) (*DeleteCustomRewardsResponse, error) {
	resp, err := c.delete("/channel_points/custom_rewards", nil, params)
	if err != nil {
		return nil, err
	}

	reward := &DeleteCustomRewardsResponse{}
	resp.HydrateResponseCommon(&reward.ResponseCommon)

	return reward, nil
}

// GetCustomRewards : Update Custom Rewards on a channel
// Required scope: channel:manage:redemptions
func (c *Client) UpdateCustomRewards(params *UpdateCustomRewardsParams) (*ChannelCustomRewardResponse, error) {
	resp, err := c.patchAsJSON("/channel_points/custom_rewards", &ManyChannelCustomRewards{}, params)
	if err != nil {
		return nil, err
	}

	rewards := &ChannelCustomRewardResponse{}
	resp.HydrateResponseCommon(&rewards.ResponseCommon)
	rewards.Data.ChannelCustomRewards = resp.Data.(*ManyChannelCustomRewards).ChannelCustomRewards

	return rewards, nil
}

// GetCustomRewards : Get Custom Rewards on a channel
// Required scope: channel:read:redemptions
func (c *Client) GetCustomRewards(params *GetCustomRewardsParams) (*ChannelCustomRewardResponse, error) {
	resp, err := c.get("/channel_points/custom_rewards", &ManyChannelCustomRewards{}, params)
	if err != nil {
		return nil, err
	}

	rewards := &ChannelCustomRewardResponse{}
	resp.HydrateResponseCommon(&rewards.ResponseCommon)
	rewards.Data.ChannelCustomRewards = resp.Data.(*ManyChannelCustomRewards).ChannelCustomRewards

	return rewards, nil
}

func (c *Client) GetCustomRewardRedemption(params *GetCustomRewardsRedemptionParams) (*RedemptionResponse, error) {
	resp, err := c.get("/channel_points/custom_rewards/redemptions", &ManyCustomRewardRedemption{}, params)
	if err != nil {
		return nil, err
	}

	rewards := &RedemptionResponse{}
	resp.HydrateResponseCommon(&rewards.ResponseCommon)
	rewards.Data.CustomRewardRedemption = resp.Data.(*ManyCustomRewardRedemption).CustomRewardRedemption

	return rewards, nil
}
