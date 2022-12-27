/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Datasource is the Schema for the datasources API
type Datasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DatasourceSpec `json:"spec,omitempty"`
}

// DatasourceSpec defines the desired state of Datasource
type DatasourceSpec struct {
	Args        []ArgsSpec        `json:"args,omitempty"`
	DataSources []DataScourceSpec `json:"datasources,omitempty"`
}

type ArgsSpec struct {
	Name      string        `json:"name"`
	Value     *string       `json:"value,omitempty"`
	ValueFrom *ValueFromCRD `json:"valueFrom,omitempty"`
}

type ValueFromCRD struct {
	SecretKeyRef *SecretRef `json:"secretKeyRef,omitempty"`
}

type SecretRef struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type DataScourceSpec struct {
	Name     string             `json:"name"`
	Enabled  bool               `json:"enabled"`
	Provider DataSourceProvider `json:"provider"`
}

type DataSourceProvider struct {
	CloudWatch  *CloudWatchProperties  `json:"cloudWatch,omitempty"`
	AppDynamics *AppDynamicProperties  `json:"appdynamics,omitempty"`
	DataDog     *DataDogProperties     `json:"datadog,omitempty"`
	Elastic     *ElasticProperties     `json:"elastic,omitempty"`
	Graphite    *GraphiteProperties    `json:"graphite,omitempty"`
	GrayLog     *Graylogproperties     `json:"graylog,omitempty"`
	NewRelic    *NewRelicProperties    `json:"newRelic,omitempty"`
	Prometheus  *PrometheusProperties  `json:"prometheus,omitempty"`
	Splunk      *SplunkProperties      `json:"splunk,omitempty"`
	StackDriver *StackDriverProperties `json:"stackDriver,omitempty"`
	SumoLogic   *SumoLogicProperties   `json:"sumologic,omitempty"`
	VmwareTanzu *VmwareTanzuProperties `json:"vmwareTanzu,omitempty"`
}

type CloudWatchProperties struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type AppDynamicProperties struct {
	ControllerHost       string `json:"controllerHost"`
	TemporaryAccessToken string `json:"temporaryAccessToken"`
}

type DataDogProperties struct {
	ApiKey         string `json:"apiKey"`
	ApplicationKey string `json:"applicationKey"`
}

type ElasticProperties struct {
	ElasticEndpoint string `json:"elasticEndpoint"`
	ElasticUsername string `json:"elasticUsername,omitempty"`
	ElasticPassword string `json:"elasticPassword,omitempty"`
	KibanaEndpoint  string `json:"kibanaEndpoint,omitempty"`
	KibanaUsername  string `json:"kibanaUsername,omitempty"`
	KibanaPassword  string `json:"kibanaPassword,omitempty"`
}

type GraphiteProperties struct {
	Endpoint string `json:"endpoint"`
}

type Graylogproperties struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

type NewRelicProperties struct {
	AccountId int32  `json:"accountId"`
	ApiKey    string `json:"apiKey"`
}

type PrometheusProperties struct {
	Endpoint string `json:"endpoint"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SplunkProperties struct {
	SplunkUrl          string `json:"splunkUrl"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	SplunkDashboardUrl string `json:"splunkDashboardUrl,omitempty"`
}

type StackDriverProperties struct {
	EncryptedKey string `json:"encryptedKey"`
}

type SumoLogicProperties struct {
	AccessId  string `json:"accessId"`
	AccessKey string `json:"accessKey"`
	Zone      string `json:"zone"`
}

type VmwareTanzuProperties struct {
	Endpoint string `json:"endpoint"`
	Email    string `json:"email"`
	ApiToken string `json:"apiToken"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatasourceList contains a list of Datasource
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datasource `json:"items"`
}
