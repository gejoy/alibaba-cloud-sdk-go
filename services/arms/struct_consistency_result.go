package arms

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

// ConsistencyResult is a nested struct in arms response
type ConsistencyResult struct {
	JudgeTimestamp                 int64                 `json:"JudgeTimestamp" xml:"JudgeTimestamp"`
	LastSlsShardWatermarkTimestamp int64                 `json:"LastSlsShardWatermarkTimestamp" xml:"LastSlsShardWatermarkTimestamp"`
	LastAgentInfoTimestamp         int64                 `json:"LastAgentInfoTimestamp" xml:"LastAgentInfoTimestamp"`
	BarrierPointsLastTimestampMap  string                `json:"BarrierPointsLastTimestampMap" xml:"BarrierPointsLastTimestampMap"`
	LastBarrierTimestamp           int64                 `json:"LastBarrierTimestamp" xml:"LastBarrierTimestamp"`
	SlsShardWatermarkBean          SlsShardWatermarkBean `json:"SlsShardWatermarkBean" xml:"SlsShardWatermarkBean"`
	AgentHeartbeatInfos            AgentHeartbeatInfos   `json:"AgentHeartbeatInfos" xml:"AgentHeartbeatInfos"`
}
