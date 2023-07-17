package schedulerx2

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

// Data is a nested struct in schedulerx2 response
type Data struct {
	Xattrs               string                   `json:"Xattrs" xml:"Xattrs"`
	JobInstanceId        int64                    `json:"JobInstanceId" xml:"JobInstanceId"`
	MonitorConfigJson    string                   `json:"MonitorConfigJson" xml:"MonitorConfigJson"`
	WfInstanceId         int64                    `json:"WfInstanceId" xml:"WfInstanceId"`
	ReadyInstanceNum     int                      `json:"ReadyInstanceNum" xml:"ReadyInstanceNum"`
	MetricsThresholdJson string                   `json:"MetricsThresholdJson" xml:"MetricsThresholdJson"`
	GroupId              string                   `json:"GroupId" xml:"GroupId"`
	AlarmJson            string                   `json:"AlarmJson" xml:"AlarmJson"`
	Description          string                   `json:"Description" xml:"Description"`
	NamespaceUid         string                   `json:"NamespaceUid" xml:"NamespaceUid"`
	JobId                int64                    `json:"JobId" xml:"JobId"`
	MaxJobs              int                      `json:"MaxJobs" xml:"MaxJobs"`
	AppKey               string                   `json:"AppKey" xml:"AppKey"`
	MaxConcurrency       int                      `json:"MaxConcurrency" xml:"MaxConcurrency"`
	CurJobs              int                      `json:"CurJobs" xml:"CurJobs"`
	WorkflowId           int64                    `json:"WorkflowId" xml:"WorkflowId"`
	RunningInstanceNum   int                      `json:"RunningInstanceNum" xml:"RunningInstanceNum"`
	AppGroupId           int64                    `json:"AppGroupId" xml:"AppGroupId"`
	AppName              string                   `json:"AppName" xml:"AppName"`
	Logs                 []string                 `json:"Logs" xml:"Logs"`
	WfInstanceDag        WfInstanceDag            `json:"WfInstanceDag" xml:"WfInstanceDag"`
	JobInstanceDetail    JobInstanceDetail        `json:"JobInstanceDetail" xml:"JobInstanceDetail"`
	WorkFlowInfo         WorkFlowInfo             `json:"WorkFlowInfo" xml:"WorkFlowInfo"`
	JobConfigInfo        JobConfigInfo            `json:"JobConfigInfo" xml:"JobConfigInfo"`
	WorkFlowNodeInfo     WorkFlowNodeInfo         `json:"WorkFlowNodeInfo" xml:"WorkFlowNodeInfo"`
	WfInstanceInfo       WfInstanceInfo           `json:"WfInstanceInfo" xml:"WfInstanceInfo"`
	WorkerInfos          []WorkerInfo             `json:"WorkerInfos" xml:"WorkerInfos"`
	AppGroups            []AppGroup               `json:"AppGroups" xml:"AppGroups"`
	Namespaces           []Namespace              `json:"Namespaces" xml:"Namespaces"`
	JobInstanceDetails   []JobInstanceDetailsItem `json:"JobInstanceDetails" xml:"JobInstanceDetails"`
	Jobs                 []Job                    `json:"Jobs" xml:"Jobs"`
	WfInstanceInfos      []WfInstanceInfosItem    `json:"WfInstanceInfos" xml:"WfInstanceInfos"`
}
