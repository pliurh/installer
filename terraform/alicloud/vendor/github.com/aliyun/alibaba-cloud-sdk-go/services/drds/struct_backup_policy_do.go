package drds

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

// BackupPolicyDO is a nested struct in drds response
type BackupPolicyDO struct {
	NextBackupActuallyTime    string `json:"NextBackupActuallyTime" xml:"NextBackupActuallyTime"`
	BackupMode                string `json:"BackupMode" xml:"BackupMode"`
	PreferredBackupPeriod     string `json:"PreferredBackupPeriod" xml:"PreferredBackupPeriod"`
	DataBackupRetentionPeriod int64  `json:"DataBackupRetentionPeriod" xml:"DataBackupRetentionPeriod"`
	BackupAppName             string `json:"BackupAppName" xml:"BackupAppName"`
	LocalLogRetentionSpace    int64  `json:"LocalLogRetentionSpace" xml:"LocalLogRetentionSpace"`
	HighSpaceUsageProtection  int64  `json:"HighSpaceUsageProtection" xml:"HighSpaceUsageProtection"`
	BackupType                string `json:"BackupType" xml:"BackupType"`
	BackupLevel               string `json:"BackupLevel" xml:"BackupLevel"`
	LocalLogRetentionHours    int64  `json:"LocalLogRetentionHours" xml:"LocalLogRetentionHours"`
	GmtCreate                 int64  `json:"GmtCreate" xml:"GmtCreate"`
	PreferredBackupTime       string `json:"PreferredBackupTime" xml:"PreferredBackupTime"`
	BackupDbName              string `json:"BackupDbName" xml:"BackupDbName"`
	BackupRetentionPeriod     int64  `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod"`
	GmtModified               int64  `json:"GmtModified" xml:"GmtModified"`
	BackupLog                 string `json:"BackupLog" xml:"BackupLog"`
	LogBackupRetentionPeriod  int64  `json:"LogBackupRetentionPeriod" xml:"LogBackupRetentionPeriod"`
	BackupPolicyMode          string `json:"BackupPolicyMode" xml:"BackupPolicyMode"`
}