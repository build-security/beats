// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreDBInstanceFromS3Input struct {
	_ struct{} `type:"structure"`

	// The amount of storage (in gigabytes) to allocate initially for the DB instance.
	// Follow the allocation rules specified in CreateDBInstance.
	//
	// Be sure to allocate enough memory for your new DB instance so that the restore
	// operation can succeed. You can also allocate additional memory for future
	// growth.
	AllocatedStorage *int64 `type:"integer"`

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the DB instance during the maintenance window. By default, minor engine
	// upgrades are not applied automatically.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone that the DB instance is created in. For information
	// about AWS Regions and Availability Zones, see Regions and Availability Zones
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html)
	// in the Amazon RDS User Guide.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS
	// Region.
	//
	// Example: us-east-1d
	//
	// Constraint: The AvailabilityZone parameter can't be specified if the DB instance
	// is a Multi-AZ deployment. The specified Availability Zone must be in the
	// same AWS Region as the current endpoint.
	AvailabilityZone *string `type:"string"`

	// The number of days for which automated backups are retained. Setting this
	// parameter to a positive number enables backups. For more information, see
	// CreateDBInstance.
	BackupRetentionPeriod *int64 `type:"integer"`

	// A value that indicates whether to copy all tags from the DB instance to snapshots
	// of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions, or for all
	// database engines. For the full list of DB instance classes, and availability
	// for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// Importing from Amazon S3 isn't supported on the db.t2.micro DB instance class.
	//
	// DBInstanceClass is a required field
	DBInstanceClass *string `type:"string" required:"true"`

	// The DB instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The name of the database to create when the DB instance is created. Follow
	// the naming rules specified in CreateDBInstance.
	DBName *string `type:"string"`

	// The name of the DB parameter group to associate with this DB instance.
	//
	// If you do not specify a value for DBParameterGroupName, then the default
	// DBParameterGroup for the specified DB engine is used.
	DBParameterGroupName *string `type:"string"`

	// A list of DB security groups to associate with this DB instance.
	//
	// Default: The default DB security group for the database engine.
	DBSecurityGroups []string `locationNameList:"DBSecurityGroupName" type:"list"`

	// A DB subnet group to associate with this DB instance.
	DBSubnetGroupName *string `type:"string"`

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `type:"boolean"`

	// The list of logs that the restored DB instance is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance.
	//
	// For more information about IAM database authentication, see IAM Database
	// Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// A value that indicates whether to enable Performance Insights for the DB
	// instance.
	//
	// For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon Relational Database Service User Guide.
	EnablePerformanceInsights *bool `type:"boolean"`

	// The name of the database engine to be used for this instance.
	//
	// Valid Values: mysql
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version number of the database engine to use. Choose the latest minor
	// version of your database engine. For information about engine versions, see
	// CreateDBInstance, or call DescribeDBEngineVersions.
	EngineVersion *string `type:"string"`

	// The amount of Provisioned IOPS (input/output operations per second) to allocate
	// initially for the DB instance. For information about valid Iops values, see
	// Amazon RDS Provisioned IOPS Storage to Improve Performance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide.
	Iops *int64 `type:"integer"`

	// The AWS KMS key identifier for an encrypted DB instance.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB instance with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB instance, then you can
	// use the KMS key alias instead of the ARN for the KM encryption key.
	//
	// If the StorageEncrypted parameter is enabled, and you do not specify a value
	// for the KmsKeyId parameter, then Amazon RDS will use your default encryption
	// key. AWS KMS creates the default encryption key for your AWS account. Your
	// AWS account has a different default encryption key for each AWS Region.
	KmsKeyId *string `type:"string"`

	// The license model for this DB instance. Use general-public-license.
	LicenseModel *string `type:"string"`

	// The password for the master user. The password can include any printable
	// ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain from 8 to 41 characters.
	MasterUserPassword *string `type:"string"`

	// The name for the master user.
	//
	// Constraints:
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	MasterUsername *string `type:"string"`

	// The interval, in seconds, between points when Enhanced Monitoring metrics
	// are collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0.
	//
	// If MonitoringRoleArn is specified, then you must also set MonitoringInterval
	// to a value other than 0.
	//
	// Valid Values: 0, 1, 5, 10, 15, 30, 60
	//
	// Default: 0
	MonitoringInterval *int64 `type:"integer"`

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics
	// to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.
	// For information on creating a monitoring role, see Setting Up and Enabling
	// Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)
	// in the Amazon RDS User Guide.
	//
	// If MonitoringInterval is set to a value other than 0, then you must supply
	// a MonitoringRoleArn value.
	MonitoringRoleArn *string `type:"string"`

	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	// If the DB instance is a Multi-AZ deployment, you can't set the AvailabilityZone
	// parameter.
	MultiAZ *bool `type:"boolean"`

	// The name of the option group to associate with this DB instance. If this
	// argument is omitted, the default option group for the specified engine is
	// used.
	OptionGroupName *string `type:"string"`

	// The AWS KMS key identifier for encryption of Performance Insights data. The
	// KMS key ID is the Amazon Resource Name (ARN), the KMS key identifier, or
	// the KMS key alias for the KMS encryption key.
	//
	// If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// RDS uses your default encryption key. AWS KMS creates the default encryption
	// key for your AWS account. Your AWS account has a different default encryption
	// key for each AWS Region.
	PerformanceInsightsKMSKeyId *string `type:"string"`

	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *int64 `type:"integer"`

	// The port number on which the database accepts connections.
	//
	// Type: Integer
	//
	// Valid Values: 1150-65535
	//
	// Default: 3306
	Port *int64 `type:"integer"`

	// The time range each day during which automated backups are created if automated
	// backups are enabled. For more information, see The Backup Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)
	// in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). For more information, see Amazon RDS Maintenance
	// Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance)
	// in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format ddd:hh24:mi-ddd:hh24:mi.
	//
	//    * Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred backup window.
	//
	//    * Must be at least 30 minutes.
	PreferredMaintenanceWindow *string `type:"string"`

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []ProcessorFeature `locationNameList:"ProcessorFeature" type:"list"`

	// A value that indicates whether the DB instance is publicly accessible.
	//
	// When the DB instance is publicly accessible, its DNS endpoint resolves to
	// the private IP address from within the DB instance's VPC, and to the public
	// IP address from outside of the DB instance's VPC. Access to the DB instance
	// is ultimately controlled by the security group it uses, and that public access
	// is not permitted if the security group assigned to the DB instance doesn't
	// permit it.
	//
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address.
	//
	// For more information, see CreateDBInstance.
	PubliclyAccessible *bool `type:"boolean"`

	// The name of your Amazon S3 bucket that contains your database backup file.
	//
	// S3BucketName is a required field
	S3BucketName *string `type:"string" required:"true"`

	// An AWS Identity and Access Management (IAM) role to allow Amazon RDS to access
	// your Amazon S3 bucket.
	//
	// S3IngestionRoleArn is a required field
	S3IngestionRoleArn *string `type:"string" required:"true"`

	// The prefix of your Amazon S3 bucket.
	S3Prefix *string `type:"string"`

	// The name of the engine of your source database.
	//
	// Valid Values: mysql
	//
	// SourceEngine is a required field
	SourceEngine *string `type:"string" required:"true"`

	// The version of the database that the backup files were created from.
	//
	// MySQL versions 5.6 and 5.7 are supported.
	//
	// Example: 5.6.40
	//
	// SourceEngineVersion is a required field
	SourceEngineVersion *string `type:"string" required:"true"`

	// A value that indicates whether the new DB instance is encrypted or not.
	StorageEncrypted *bool `type:"boolean"`

	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified; otherwise gp2
	StorageType *string `type:"string"`

	// A list of tags to associate with this DB instance. For more information,
	// see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A value that indicates whether the DB instance class of the DB instance uses
	// its default processor features.
	UseDefaultProcessorFeatures *bool `type:"boolean"`

	// A list of VPC security groups to associate with this DB instance.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBInstanceFromS3Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBInstanceFromS3Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBInstanceFromS3Input"}

	if s.DBInstanceClass == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceClass"))
	}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if s.S3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketName"))
	}

	if s.S3IngestionRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3IngestionRoleArn"))
	}

	if s.SourceEngine == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEngine"))
	}

	if s.SourceEngineVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEngineVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreDBInstanceFromS3Output struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RestoreDBInstanceFromS3Output) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBInstanceFromS3 = "RestoreDBInstanceFromS3"

// RestoreDBInstanceFromS3Request returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Amazon Relational Database Service (Amazon RDS) supports importing MySQL
// databases by using backup files. You can create a backup of your on-premises
// database, store it on Amazon Simple Storage Service (Amazon S3), and then
// restore the backup file onto a new Amazon RDS DB instance running MySQL.
// For more information, see Importing Data into an Amazon RDS MySQL DB Instance
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html)
// in the Amazon RDS User Guide.
//
//    // Example sending a request using RestoreDBInstanceFromS3Request.
//    req := client.RestoreDBInstanceFromS3Request(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceFromS3
func (c *Client) RestoreDBInstanceFromS3Request(input *RestoreDBInstanceFromS3Input) RestoreDBInstanceFromS3Request {
	op := &aws.Operation{
		Name:       opRestoreDBInstanceFromS3,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBInstanceFromS3Input{}
	}

	req := c.newRequest(op, input, &RestoreDBInstanceFromS3Output{})

	return RestoreDBInstanceFromS3Request{Request: req, Input: input, Copy: c.RestoreDBInstanceFromS3Request}
}

// RestoreDBInstanceFromS3Request is the request type for the
// RestoreDBInstanceFromS3 API operation.
type RestoreDBInstanceFromS3Request struct {
	*aws.Request
	Input *RestoreDBInstanceFromS3Input
	Copy  func(*RestoreDBInstanceFromS3Input) RestoreDBInstanceFromS3Request
}

// Send marshals and sends the RestoreDBInstanceFromS3 API request.
func (r RestoreDBInstanceFromS3Request) Send(ctx context.Context) (*RestoreDBInstanceFromS3Response, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBInstanceFromS3Response{
		RestoreDBInstanceFromS3Output: r.Request.Data.(*RestoreDBInstanceFromS3Output),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBInstanceFromS3Response is the response type for the
// RestoreDBInstanceFromS3 API operation.
type RestoreDBInstanceFromS3Response struct {
	*RestoreDBInstanceFromS3Output

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBInstanceFromS3 request.
func (r *RestoreDBInstanceFromS3Response) SDKResponseMetdata() *aws.Response {
	return r.response
}