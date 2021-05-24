package ccc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Inbound is a nested struct in ccc response
type Inbound struct {
	AverageHoldTime              float64 `json:"AverageHoldTime" xml:"AverageHoldTime"`
	MaxAbandonedInQueueTime      int64   `json:"MaxAbandonedInQueueTime" xml:"MaxAbandonedInQueueTime"`
	AverageAbandonTime           float64 `json:"AverageAbandonTime" xml:"AverageAbandonTime"`
	CallsConsulted               int64   `json:"CallsConsulted" xml:"CallsConsulted"`
	MaxAbandonedInIVRTime        int64   `json:"MaxAbandonedInIVRTime" xml:"MaxAbandonedInIVRTime"`
	CallsAbandonedInRinging      int64   `json:"CallsAbandonedInRinging" xml:"CallsAbandonedInRinging"`
	TotalRingTime                int64   `json:"TotalRingTime" xml:"TotalRingTime"`
	AbandonRate                  float64 `json:"AbandonRate" xml:"AbandonRate"`
	MaxRingTime                  int64   `json:"MaxRingTime" xml:"MaxRingTime"`
	CallsQueued                  int64   `json:"CallsQueued" xml:"CallsQueued"`
	MaxWaitTime                  int64   `json:"MaxWaitTime" xml:"MaxWaitTime"`
	AverageTalkTime              float64 `json:"AverageTalkTime" xml:"AverageTalkTime"`
	TotalWaitTime                int64   `json:"TotalWaitTime" xml:"TotalWaitTime"`
	CallsAbandonedInQueue        int64   `json:"CallsAbandonedInQueue" xml:"CallsAbandonedInQueue"`
	MaxAbandonedInRingTime       int64   `json:"MaxAbandonedInRingTime" xml:"MaxAbandonedInRingTime"`
	MaxWorkTime                  int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	TotalWorkTime                int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
	SatisfactionSurveysOffered   int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	CallsHold                    int64   `json:"CallsHold" xml:"CallsHold"`
	MaxAbandonTime               int64   `json:"MaxAbandonTime" xml:"MaxAbandonTime"`
	TotalHoldTime                int64   `json:"TotalHoldTime" xml:"TotalHoldTime"`
	CallsAbandoned               int64   `json:"CallsAbandoned" xml:"CallsAbandoned"`
	MaxHoldTime                  int64   `json:"MaxHoldTime" xml:"MaxHoldTime"`
	SatisfactionSurveysResponded int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	ServiceLevel20               float64 `json:"ServiceLevel20" xml:"ServiceLevel20"`
	AverageAbandonedInIVRTime    float64 `json:"AverageAbandonedInIVRTime" xml:"AverageAbandonedInIVRTime"`
	HandleRate                   float64 `json:"HandleRate" xml:"HandleRate"`
	TotalAbandonedInIVRTime      int64   `json:"TotalAbandonedInIVRTime" xml:"TotalAbandonedInIVRTime"`
	AverageRingTime              float64 `json:"AverageRingTime" xml:"AverageRingTime"`
	SatisfactionIndex            float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	AverageWaitTime              float64 `json:"AverageWaitTime" xml:"AverageWaitTime"`
	AverageAbandonedInQueueTime  float64 `json:"AverageAbandonedInQueueTime" xml:"AverageAbandonedInQueueTime"`
	MaxTalkTime                  int64   `json:"MaxTalkTime" xml:"MaxTalkTime"`
	AbandonedRate                float64 `json:"AbandonedRate" xml:"AbandonedRate"`
	CallsAbandonedInIVR          int64   `json:"CallsAbandonedInIVR" xml:"CallsAbandonedInIVR"`
	TotalAbandonedInQueueTime    int64   `json:"TotalAbandonedInQueueTime" xml:"TotalAbandonedInQueueTime"`
	TotalTalkTime                int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	CallsOffered                 int64   `json:"CallsOffered" xml:"CallsOffered"`
	TotalAbandonTime             int64   `json:"TotalAbandonTime" xml:"TotalAbandonTime"`
	AverageAbandonedInRingTime   float64 `json:"AverageAbandonedInRingTime" xml:"AverageAbandonedInRingTime"`
	CallsRinged                  int64   `json:"CallsRinged" xml:"CallsRinged"`
	CallsHandled                 int64   `json:"CallsHandled" xml:"CallsHandled"`
	TotalAbandonedInRingTime     int64   `json:"TotalAbandonedInRingTime" xml:"TotalAbandonedInRingTime"`
	CallsTransferred             int64   `json:"CallsTransferred" xml:"CallsTransferred"`
	AverageWorkTime              float64 `json:"AverageWorkTime" xml:"AverageWorkTime"`
}
