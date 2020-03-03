package outboundbot

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

// JobGroup is a nested struct in outboundbot response
type JobGroup struct {
	ScriptName          string   `json:"ScriptName" xml:"ScriptName"`
	JobGroupId          string   `json:"JobGroupId" xml:"JobGroupId"`
	JobGroupName        string   `json:"JobGroupName" xml:"JobGroupName"`
	CreationTime        int64    `json:"CreationTime" xml:"CreationTime"`
	ScenarioId          string   `json:"ScenarioId" xml:"ScenarioId"`
	ScriptId            string   `json:"ScriptId" xml:"ScriptId"`
	JobGroupDescription string   `json:"JobGroupDescription" xml:"JobGroupDescription"`
	JobFilePath         string   `json:"JobFilePath" xml:"JobFilePath"`
	ScenarioName        string   `json:"ScenarioName" xml:"ScenarioName"`
	CallingNumbers      []string `json:"CallingNumbers" xml:"CallingNumbers"`
	Strategy            Strategy `json:"Strategy" xml:"Strategy"`
	Progress            Progress `json:"Progress" xml:"Progress"`
}