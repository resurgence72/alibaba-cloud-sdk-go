package config

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

// ConfigRule is a nested struct in config response
type ConfigRule struct {
	RiskLevel                  int                        `json:"RiskLevel" xml:"RiskLevel"`
	ExcludeResourceIdsScope    string                     `json:"ExcludeResourceIdsScope" xml:"ExcludeResourceIdsScope"`
	CreateTimestamp            int64                      `json:"CreateTimestamp" xml:"CreateTimestamp"`
	ResourceGroupIdsScope      string                     `json:"ResourceGroupIdsScope" xml:"ResourceGroupIdsScope"`
	SourceOwner                string                     `json:"SourceOwner" xml:"SourceOwner"`
	ModifiedTimestamp          int64                      `json:"ModifiedTimestamp" xml:"ModifiedTimestamp"`
	TagValueScope              string                     `json:"TagValueScope" xml:"TagValueScope"`
	AutomationType             string                     `json:"AutomationType" xml:"AutomationType"`
	TagKeyScope                string                     `json:"TagKeyScope" xml:"TagKeyScope"`
	RegionIdsScope             string                     `json:"RegionIdsScope" xml:"RegionIdsScope"`
	ConfigRuleState            string                     `json:"ConfigRuleState" xml:"ConfigRuleState"`
	Description                string                     `json:"Description" xml:"Description"`
	ConfigRuleTriggerTypes     string                     `json:"ConfigRuleTriggerTypes" xml:"ConfigRuleTriggerTypes"`
	MaximumExecutionFrequency  string                     `json:"MaximumExecutionFrequency" xml:"MaximumExecutionFrequency"`
	ConfigRuleName             string                     `json:"ConfigRuleName" xml:"ConfigRuleName"`
	ConfigRuleId               string                     `json:"ConfigRuleId" xml:"ConfigRuleId"`
	ConfigRuleArn              string                     `json:"ConfigRuleArn" xml:"ConfigRuleArn"`
	InputParameters            map[string]interface{}     `json:"InputParameters" xml:"InputParameters"`
	AccountId                  int64                      `json:"AccountId" xml:"AccountId"`
	SourceIdentifier           string                     `json:"SourceIdentifier" xml:"SourceIdentifier"`
	ResourceTypesScope         string                     `json:"ResourceTypesScope" xml:"ResourceTypesScope"`
	CreateBy                   CreateBy                   `json:"CreateBy" xml:"CreateBy"`
	ConfigRuleEvaluationStatus ConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus" xml:"ConfigRuleEvaluationStatus"`
	Source                     Source                     `json:"Source" xml:"Source"`
	Compliance                 Compliance                 `json:"Compliance" xml:"Compliance"`
	Scope                      Scope                      `json:"Scope" xml:"Scope"`
	ManagedRule                ManagedRule                `json:"ManagedRule" xml:"ManagedRule"`
}
