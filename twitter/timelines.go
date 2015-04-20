package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

type TimelineService struct {
	sling *sling.Sling
}

// NewTimelineService returns a new TimelineService for accessing timeline
// statuses API endpoints.
func NewTimelineService(sling *sling.Sling) *TimelineService {
	return &TimelineService{
		sling: sling.Path("statuses/"),
	}
}

type UserTimelineParams struct {
	UserId             int64  `url:"user_id,omitempty"`
	ScreenName         string `url:"screen_name,omitempty"`
	Count              int    `url:"count,omitempty"`
	SinceId            int64  `url:"since_id,omitempty"`
	MaxId              int64  `url:"max_id,omitempty"`
	TrimUser           *bool  `url:"trim_user,omitempty"`
	ExcludeReplies     *bool  `url:"exclude_replies,omitempty"`
	ContributorDetails *bool  `url:"contributor_details,omitempty"`
	IncludeRetweets    *bool  `url:"include_rts,omitempty"`
}

// UserTimeline returns the user timeline of recent tweets by the specified user.
// https://dev.twitter.com/rest/reference/get/statuses/user_timeline
func (s *TimelineService) UserTimeline(params *UserTimelineParams) ([]Tweet, *http.Response, error) {
	tweets := new([]Tweet)
	resp, err := s.sling.New().Get("user_timeline.json").QueryStruct(params).Receive(tweets)
	return *tweets, resp, err
}

type HomeTimelineParams struct {
	Count              int   `url:"count,omitempty"`
	SinceId            int64 `url:"since_id,omitempty"`
	MaxId              int64 `url:"max_id,omitempty"`
	TrimUser           *bool `url:"trim_user,omitempty"`
	ExcludeReplies     *bool `url:"exclude_replies,omitempty"`
	ContributorDetails *bool `url:"contributor_details,omitempty"`
	IncludeEntities    *bool `url:"include_entities,omitempty"`
}

// HomeTimeline returns recent Tweets and retweets from the user and those users
// they follow.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/get/statuses/home_timeline
func (s *TimelineService) HomeTimeline(params *HomeTimelineParams) ([]Tweet, *http.Response, error) {
	tweets := new([]Tweet)
	resp, err := s.sling.New().Get("home_timeline.json").QueryStruct(params).Receive(tweets)
	return *tweets, resp, err
}

type MentionTimelineParams struct {
	Count              int   `url:"count,omitempty"`
	SinceId            int64 `url:"since_id,omitempty"`
	MaxId              int64 `url:"max_id,omitempty"`
	TrimUser           *bool `url:"trim_user,omitempty"`
	ContributorDetails *bool `url:"contributor_details,omitempty"`
	IncludeEntities    *bool `url:"include_entities,omitempty"`
}

// MentionTimeline returns the most recent mentions of the authenticated user.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/get/statuses/mentions_timeline
func (s *TimelineService) MentionTimeline(params *MentionTimelineParams) ([]Tweet, *http.Response, error) {
	tweets := new([]Tweet)
	resp, err := s.sling.New().Get("mentions_timeline.json").QueryStruct(params).Receive(tweets)
	return *tweets, resp, err
}

type RetweetsOfMeTimelineParams struct {
	Count               int   `url:"count,omitempty"`
	SinceId             int64 `url:"since_id,omitempty"`
	MaxId               int64 `url:"max_id,omitempty"`
	TrimUser            *bool `url:"trim_user,omitempty"`
	IncludeEntities     *bool `url:"include_entities,omitempty"`
	IncludeUserEntities *bool `url:"include_user_entities"`
}

// RetweetsOfMeTimeline returns the most recent Tweets by the authenticated
// user that have been retweeted by others.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/get/statuses/retweets_of_me
func (s *TimelineService) RetweetsOfMeTimeline(params *RetweetsOfMeTimelineParams) ([]Tweet, *http.Response, error) {
	tweets := new([]Tweet)
	resp, err := s.sling.New().Get("retweets_of_me.json").QueryStruct(params).Receive(tweets)
	return *tweets, resp, err
}
