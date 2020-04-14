// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package output

import (
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/common"
	"github.com/banzaicloud/operator-tools/pkg/secret"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureStorage) DeepCopyInto(out *AzureStorage) {
	*out = *in
	if in.AzureStorageAccount != nil {
		in, out := &in.AzureStorageAccount, &out.AzureStorageAccount
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AzureStorageAccessKey != nil {
		in, out := &in.AzureStorageAccessKey, &out.AzureStorageAccessKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureStorage.
func (in *AzureStorage) DeepCopy() *AzureStorage {
	if in == nil {
		return nil
	}
	out := new(AzureStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Buffer) DeepCopyInto(out *Buffer) {
	*out = *in
	if in.RetryForever != nil {
		in, out := &in.RetryForever, &out.RetryForever
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Buffer.
func (in *Buffer) DeepCopy() *Buffer {
	if in == nil {
		return nil
	}
	out := new(Buffer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudWatchOutput) DeepCopyInto(out *CloudWatchOutput) {
	*out = *in
	if in.AwsAccessKey != nil {
		in, out := &in.AwsAccessKey, &out.AwsAccessKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AwsSecretKey != nil {
		in, out := &in.AwsSecretKey, &out.AwsSecretKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudWatchOutput.
func (in *CloudWatchOutput) DeepCopy() *CloudWatchOutput {
	if in == nil {
		return nil
	}
	out := new(CloudWatchOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOutput) DeepCopyInto(out *ElasticsearchOutput) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SslVerify != nil {
		in, out := &in.SslVerify, &out.SslVerify
		*out = new(bool)
		**out = **in
	}
	if in.UtcIndex != nil {
		in, out := &in.UtcIndex, &out.UtcIndex
		*out = new(bool)
		**out = **in
	}
	if in.TemplateFile != nil {
		in, out := &in.TemplateFile, &out.TemplateFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexDatePattern != nil {
		in, out := &in.IndexDatePattern, &out.IndexDatePattern
		*out = new(string)
		**out = **in
	}
	if in.ApplicationName != nil {
		in, out := &in.ApplicationName, &out.ApplicationName
		*out = new(string)
		**out = **in
	}
	if in.FailOnPuttingTemplateRetryExceed != nil {
		in, out := &in.FailOnPuttingTemplateRetryExceed, &out.FailOnPuttingTemplateRetryExceed
		*out = new(bool)
		**out = **in
	}
	if in.ReloadConnections != nil {
		in, out := &in.ReloadConnections, &out.ReloadConnections
		*out = new(bool)
		**out = **in
	}
	if in.VerifyEsVersionAtStartup != nil {
		in, out := &in.VerifyEsVersionAtStartup, &out.VerifyEsVersionAtStartup
		*out = new(bool)
		**out = **in
	}
	if in.ExceptionBackup != nil {
		in, out := &in.ExceptionBackup, &out.ExceptionBackup
		*out = new(bool)
		**out = **in
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOutput.
func (in *ElasticsearchOutput) DeepCopy() *ElasticsearchOutput {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileOutputConfig) DeepCopyInto(out *FileOutputConfig) {
	*out = *in
	if in.AddPathSuffix != nil {
		in, out := &in.AddPathSuffix, &out.AddPathSuffix
		*out = new(bool)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileOutputConfig.
func (in *FileOutputConfig) DeepCopy() *FileOutputConfig {
	if in == nil {
		return nil
	}
	out := new(FileOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdServer) DeepCopyInto(out *FluentdServer) {
	*out = *in
	if in.SharedKey != nil {
		in, out := &in.SharedKey, &out.SharedKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdServer.
func (in *FluentdServer) DeepCopy() *FluentdServer {
	if in == nil {
		return nil
	}
	out := new(FluentdServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Format) DeepCopyInto(out *Format) {
	*out = *in
	if in.AddNewline != nil {
		in, out := &in.AddNewline, &out.AddNewline
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Format.
func (in *Format) DeepCopy() *Format {
	if in == nil {
		return nil
	}
	out := new(Format)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardOutput) DeepCopyInto(out *ForwardOutput) {
	*out = *in
	if in.FluentdServers != nil {
		in, out := &in.FluentdServers, &out.FluentdServers
		*out = make([]FluentdServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TlsCertPath != nil {
		in, out := &in.TlsCertPath, &out.TlsCertPath
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.TlsClientCertPath != nil {
		in, out := &in.TlsClientCertPath, &out.TlsClientCertPath
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.TlsClientPrivateKeyPath != nil {
		in, out := &in.TlsClientPrivateKeyPath, &out.TlsClientPrivateKeyPath
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.TlsClientPrivateKeyPassphrase != nil {
		in, out := &in.TlsClientPrivateKeyPassphrase, &out.TlsClientPrivateKeyPassphrase
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(common.Security)
		**out = **in
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardOutput.
func (in *ForwardOutput) DeepCopy() *ForwardOutput {
	if in == nil {
		return nil
	}
	out := new(ForwardOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSOutput) DeepCopyInto(out *GCSOutput) {
	*out = *in
	if in.CredentialsJson != nil {
		in, out := &in.CredentialsJson, &out.CredentialsJson
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectMetadata != nil {
		in, out := &in.ObjectMetadata, &out.ObjectMetadata
		*out = make([]ObjectMetadata, len(*in))
		copy(*out, *in)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSOutput.
func (in *GCSOutput) DeepCopy() *GCSOutput {
	if in == nil {
		return nil
	}
	out := new(GCSOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaOutputConfig) DeepCopyInto(out *KafkaOutputConfig) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HeadersFromRecord != nil {
		in, out := &in.HeadersFromRecord, &out.HeadersFromRecord
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSLCACert != nil {
		in, out := &in.SSLCACert, &out.SSLCACert
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLClientCert != nil {
		in, out := &in.SSLClientCert, &out.SSLClientCert
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLClientCertChain != nil {
		in, out := &in.SSLClientCertChain, &out.SSLClientCertChain
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLClientCertKey != nil {
		in, out := &in.SSLClientCertKey, &out.SSLClientCertKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLVerifyHostname != nil {
		in, out := &in.SSLVerifyHostname, &out.SSLVerifyHostname
		*out = new(bool)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaOutputConfig.
func (in *KafkaOutputConfig) DeepCopy() *KafkaOutputConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamAssumeRoleCredentials) DeepCopyInto(out *KinesisStreamAssumeRoleCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamAssumeRoleCredentials.
func (in *KinesisStreamAssumeRoleCredentials) DeepCopy() *KinesisStreamAssumeRoleCredentials {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamAssumeRoleCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamOutputConfig) DeepCopyInto(out *KinesisStreamOutputConfig) {
	*out = *in
	if in.AWSKeyId != nil {
		in, out := &in.AWSKeyId, &out.AWSKeyId
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSSECKey != nil {
		in, out := &in.AWSSECKey, &out.AWSSECKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSSESToken != nil {
		in, out := &in.AWSSESToken, &out.AWSSESToken
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AssumeRoleCredentials != nil {
		in, out := &in.AssumeRoleCredentials, &out.AssumeRoleCredentials
		*out = new(KinesisStreamAssumeRoleCredentials)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamOutputConfig.
func (in *KinesisStreamOutputConfig) DeepCopy() *KinesisStreamOutputConfig {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogZOutput) DeepCopyInto(out *LogZOutput) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogZOutput.
func (in *LogZOutput) DeepCopy() *LogZOutput {
	if in == nil {
		return nil
	}
	out := new(LogZOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiOutput) DeepCopyInto(out *LokiOutput) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(Label, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtractKubernetesLabels != nil {
		in, out := &in.ExtractKubernetesLabels, &out.ExtractKubernetesLabels
		*out = new(bool)
		**out = **in
	}
	if in.RemoveKeys != nil {
		in, out := &in.RemoveKeys, &out.RemoveKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiOutput.
func (in *LokiOutput) DeepCopy() *LokiOutput {
	if in == nil {
		return nil
	}
	out := new(LokiOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewRelicOutputConfig) DeepCopyInto(out *NewRelicOutputConfig) {
	*out = *in
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.LicenseKey != nil {
		in, out := &in.LicenseKey, &out.LicenseKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewRelicOutputConfig.
func (in *NewRelicOutputConfig) DeepCopy() *NewRelicOutputConfig {
	if in == nil {
		return nil
	}
	out := new(NewRelicOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NullOutputConfig) DeepCopyInto(out *NullOutputConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NullOutputConfig.
func (in *NullOutputConfig) DeepCopy() *NullOutputConfig {
	if in == nil {
		return nil
	}
	out := new(NullOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSSOutput) DeepCopyInto(out *OSSOutput) {
	*out = *in
	if in.AccessKeyId != nil {
		in, out := &in.AccessKeyId, &out.AccessKeyId
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AaccessKeySecret != nil {
		in, out := &in.AaccessKeySecret, &out.AaccessKeySecret
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSSOutput.
func (in *OSSOutput) DeepCopy() *OSSOutput {
	if in == nil {
		return nil
	}
	out := new(OSSOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AssumeRoleCredentials) DeepCopyInto(out *S3AssumeRoleCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AssumeRoleCredentials.
func (in *S3AssumeRoleCredentials) DeepCopy() *S3AssumeRoleCredentials {
	if in == nil {
		return nil
	}
	out := new(S3AssumeRoleCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3InstanceProfileCredentials) DeepCopyInto(out *S3InstanceProfileCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3InstanceProfileCredentials.
func (in *S3InstanceProfileCredentials) DeepCopy() *S3InstanceProfileCredentials {
	if in == nil {
		return nil
	}
	out := new(S3InstanceProfileCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutputConfig) DeepCopyInto(out *S3OutputConfig) {
	*out = *in
	if in.AwsAccessKey != nil {
		in, out := &in.AwsAccessKey, &out.AwsAccessKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.AwsSecretKey != nil {
		in, out := &in.AwsSecretKey, &out.AwsSecretKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.AssumeRoleCredentials != nil {
		in, out := &in.AssumeRoleCredentials, &out.AssumeRoleCredentials
		*out = new(S3AssumeRoleCredentials)
		**out = **in
	}
	if in.InstanceProfileCredentials != nil {
		in, out := &in.InstanceProfileCredentials, &out.InstanceProfileCredentials
		*out = new(S3InstanceProfileCredentials)
		**out = **in
	}
	if in.SharedCredentials != nil {
		in, out := &in.SharedCredentials, &out.SharedCredentials
		*out = new(S3SharedCredentials)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutputConfig.
func (in *S3OutputConfig) DeepCopy() *S3OutputConfig {
	if in == nil {
		return nil
	}
	out := new(S3OutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3SharedCredentials) DeepCopyInto(out *S3SharedCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3SharedCredentials.
func (in *S3SharedCredentials) DeepCopy() *S3SharedCredentials {
	if in == nil {
		return nil
	}
	out := new(S3SharedCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkHecOutput) DeepCopyInto(out *SplunkHecOutput) {
	*out = *in
	if in.HecToken != nil {
		in, out := &in.HecToken, &out.HecToken
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricsFromEvent != nil {
		in, out := &in.MetricsFromEvent, &out.MetricsFromEvent
		*out = new(bool)
		**out = **in
	}
	if in.CoerceToUtf8 != nil {
		in, out := &in.CoerceToUtf8, &out.CoerceToUtf8
		*out = new(bool)
		**out = **in
	}
	if in.InsecureSSL != nil {
		in, out := &in.InsecureSSL, &out.InsecureSSL
		*out = new(bool)
		**out = **in
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Buffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkHecOutput.
func (in *SplunkHecOutput) DeepCopy() *SplunkHecOutput {
	if in == nil {
		return nil
	}
	out := new(SplunkHecOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumologicOutput) DeepCopyInto(out *SumologicOutput) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumologicOutput.
func (in *SumologicOutput) DeepCopy() *SumologicOutput {
	if in == nil {
		return nil
	}
	out := new(SumologicOutput)
	in.DeepCopyInto(out)
	return out
}
