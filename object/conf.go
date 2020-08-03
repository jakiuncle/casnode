// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

var (
	DefaultPageNum             = 20
	DefaultHomePageNum         = 50
	DefaultNotificationPageNum = 10
	DefaultBalancePageNum      = 25
	UserNamingRestrictions     = true
	HomePageNodeNum            = 8
	TopicThanksCost            = 15
	ReplyThanksCost            = 10
	CreateTopicCost            = 20
	CreateReplyCost            = 5
	ReceiveReplyBonus          = 5
	MaxDailyCheckinBonus       = 20
	LatestNodeNum              = 20
	HotNodeNum                 = 15
	HotTopicNum                = 10
	TopicEditableTime          = 10.0 //minutes
	ReplyEditableTime          = 10.0 //minutes
	ReplyDeletableTime         = 5.0  //minutes
	NodeHitRecordExpiredTime   = 1    //month
	TopicHitRecordExpiredTime  = 1    //day
	ValidateCodeExpiredTime    = 20   //minutes
	UseOAuthProxy              = false

	DefaultCronJobs = []*CronJob{
		{
			Id:       "updateExpiredData",
			BumpTime: "0:0",
			State:    "active",
		},
		{
			Id:       "updateHotInfo",
			BumpTime: "*/1:*",
			State:    "active",
		},
		{
			Id:       "expireValidateCode",
			BumpTime: "*/1:*",
			State:    "active",
		},
	}
	DefaultCronUpdates = []*UpdateJob{
		{
			Id:    "expireData",
			JobId: "updateExpiredData",
			State: "active",
		},
		{
			Id:    "hotInfo",
			JobId: "updateHotInfo",
			State: "active",
		},
		{
			Id:    "expireValidateCode",
			JobId: "expireValidateCode",
			State: "active",
		},
	}
)
