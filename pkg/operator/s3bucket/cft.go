// >>>>>>> DO NOT EDIT THIS FILE <<<<<<<<<<
// This file is autogenerated via `aws-operator-codegen process`
// If you'd like the change anything about this file make edits to the .templ
// file in the pkg/codegen/assets directory.

package s3bucket

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsV1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	"github.com/christopherhein/aws-operator/pkg/config"
	"github.com/christopherhein/aws-operator/pkg/helpers"
)

// New generates a new object
func New(config *config.Config, s3bucket *awsV1alpha1.S3Bucket, topicARN string) *Cloudformation {
	return &Cloudformation{
		S3Bucket: s3bucket,
		config:					config,
    topicARN:       topicARN,
	}
}

// Cloudformation defines the s3bucket cfts
type Cloudformation struct {
	config         *config.Config
	S3Bucket *awsV1alpha1.S3Bucket
  topicARN       string
}

// GetOutputs return the stack outputs from the DescribeStacks call
func (s *Cloudformation) GetOutputs() (map[string]string, error) {
	outputs := map[string]string{}
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DescribeStacksInput{
		StackName:   aws.String(s.S3Bucket.Name),
	}

	output, err := svc.DescribeStacks(&stackInputs)
	if err != nil {
		return nil, err
	}
	// Not sure if this is even possible
	if len(output.Stacks) != 1 {
		return nil, errors.New("no stacks returned with that stack name")
	}

	for _, out := range output.Stacks[0].Outputs {
		outputs[*out.OutputKey] = *out.OutputValue
	}

	return outputs, err
}

// CreateStack will create the stack with the supplied params
func (s *Cloudformation) CreateStack() (output *cloudformation.CreateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	cftemplate := helpers.GetCloudFormationTemplate(s.config, "s3bucket", s.S3Bucket.Spec.CloudFormationTemplateName, s.S3Bucket.Spec.CloudFormationTemplateNamespace)

	stackInputs := cloudformation.CreateStackInput{
		StackName:   aws.String(s.S3Bucket.Name),
		TemplateURL: aws.String(cftemplate),
		NotificationARNs: []*string{
			aws.String(s.topicARN),
		},
	}

	resourceName := helpers.CreateParam("ResourceName", s.S3Bucket.Name)
	resourceVersion := helpers.CreateParam("ResourceVersion", s.S3Bucket.ResourceVersion)
	namespace := helpers.CreateParam("Namespace", s.S3Bucket.Namespace)
	clusterName := helpers.CreateParam("ClusterName", s.config.ClusterName)
	bucketName := helpers.CreateParam("BucketName", helpers.Stringify(s.S3Bucket.Spec.BucketName))
	versioning := helpers.CreateParam("EnableVersioning", helpers.Stringify(s.S3Bucket.Spec.Versioning))
  loggingenabled := helpers.CreateParam("EnableLogging", helpers.Stringify(s.S3Bucket.Spec.Logging.Enabled))
  loggingprefix := helpers.CreateParam("LoggingPrefix", helpers.Stringify(s.S3Bucket.Spec.Logging.Prefix))

	parameters := []*cloudformation.Parameter{}
	parameters = append(parameters, resourceName)
	parameters = append(parameters, resourceVersion)
	parameters = append(parameters, namespace)
	parameters = append(parameters, clusterName)
	parameters = append(parameters, bucketName)
	parameters = append(parameters, versioning)
	parameters = append(parameters, loggingenabled)
	parameters = append(parameters, loggingprefix)

	stackInputs.SetParameters(parameters)

	resourceNameTag := helpers.CreateTag("ResourceName", s.S3Bucket.Name)
	resourceVersionTag := helpers.CreateTag("ResourceVersion", s.S3Bucket.ResourceVersion)
	namespaceTag := helpers.CreateTag("Namespace", s.S3Bucket.Namespace)
	clusterNameTag := helpers.CreateTag("ClusterName", s.config.ClusterName)

	tags := []*cloudformation.Tag{}
	tags = append(tags, resourceNameTag)
	tags = append(tags, resourceVersionTag)
	tags = append(tags, namespaceTag)
	tags = append(tags, clusterNameTag)

	stackInputs.SetTags(tags)

  output, err = svc.CreateStack(&stackInputs)
	return
}

// UpdateStack will update the existing stack
func (s *Cloudformation) UpdateStack(updated *awsV1alpha1.S3Bucket) (output *cloudformation.UpdateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	cftemplate := helpers.GetCloudFormationTemplate(s.config, "s3bucket", updated.Spec.CloudFormationTemplateName, updated.Spec.CloudFormationTemplateNamespace)

	stackInputs := cloudformation.UpdateStackInput{
		StackName:   aws.String(s.S3Bucket.Name),
		TemplateURL: aws.String(cftemplate),
		NotificationARNs: []*string{
			aws.String(s.topicARN),
		},
	}

	resourceName := helpers.CreateParam("ResourceName", s.S3Bucket.Name)
	resourceVersion := helpers.CreateParam("ResourceVersion", s.S3Bucket.ResourceVersion)
	namespace := helpers.CreateParam("Namespace", s.S3Bucket.Namespace)
	clusterName := helpers.CreateParam("ClusterName", s.config.ClusterName)
	bucketName := helpers.CreateParam("BucketName", helpers.Stringify(updated.Spec.BucketName))
	versioning := helpers.CreateParam("EnableVersioning", helpers.Stringify(updated.Spec.Versioning))
	loggingenabled := helpers.CreateParam("EnableLogging", helpers.Stringify(updated.Spec.Logging.Enabled))
	loggingprefix := helpers.CreateParam("LoggingPrefix", helpers.Stringify(updated.Spec.Logging.Prefix))

	parameters := []*cloudformation.Parameter{}
	parameters = append(parameters, resourceName)
	parameters = append(parameters, resourceVersion)
	parameters = append(parameters, namespace)
	parameters = append(parameters, clusterName)
	parameters = append(parameters, bucketName)
	parameters = append(parameters, versioning)
	parameters = append(parameters, loggingenabled)
	parameters = append(parameters, loggingprefix)

	stackInputs.SetParameters(parameters)

	resourceNameTag := helpers.CreateTag("ResourceName", s.S3Bucket.Name)
	resourceVersionTag := helpers.CreateTag("ResourceVersion", s.S3Bucket.ResourceVersion)
	namespaceTag := helpers.CreateTag("Namespace", s.S3Bucket.Namespace)
	clusterNameTag := helpers.CreateTag("ClusterName", s.config.ClusterName)

	tags := []*cloudformation.Tag{}
	tags = append(tags, resourceNameTag)
	tags = append(tags, resourceVersionTag)
	tags = append(tags, namespaceTag)
	tags = append(tags, clusterNameTag)

	stackInputs.SetTags(tags)

  output, err = svc.UpdateStack(&stackInputs)
	return
}

// DeleteStack will delete the stack
func (s *Cloudformation) DeleteStack() (err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DeleteStackInput{}
	stackInputs.SetStackName(s.S3Bucket.Name)

  _, err = svc.DeleteStack(&stackInputs)
	return
}