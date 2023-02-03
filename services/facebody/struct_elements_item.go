package facebody

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

// ElementsItem is a nested struct in facebody response
type ElementsItem struct {
	Category    string        `json:"Category" xml:"Category"`
	TemplateId  string        `json:"TemplateId" xml:"TemplateId"`
	TemplateURL string        `json:"TemplateURL" xml:"TemplateURL"`
	UpdateTime  string        `json:"UpdateTime" xml:"UpdateTime"`
	Confidence  float64       `json:"Confidence" xml:"Confidence"`
	UserId      string        `json:"UserId" xml:"UserId"`
	CreateTime  string        `json:"CreateTime" xml:"CreateTime"`
	ImageURL    string        `json:"ImageURL" xml:"ImageURL"`
	FaceNumber  int64         `json:"FaceNumber" xml:"FaceNumber"`
	Box         []float64     `json:"Box" xml:"Box"`
	Results     []ResultsItem `json:"Results" xml:"Results"`
}
