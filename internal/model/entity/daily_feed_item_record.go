// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-08-06 17:07:07
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DailyFeedItemRecord is the golang structure for table daily_feed_item_record.
type DailyFeedItemRecord struct {
	ChannelId string      `json:"channelId" ` //
	ItemId    string      `json:"itemId"    ` //
	PubDate   *gtime.Time `json:"pubDate"   ` //
}
