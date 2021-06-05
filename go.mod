module github.com/evanboyle/singleBinaryAutomationAPI

go 1.16

require (
	github.com/pulumi/pulumi/auto/v3 v3.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi/sdk/v3 v3.3.1
)

replace (
	github.com/pulumi/pulumi/auto/v3 => ../../pulumi/pulumi/auto
	github.com/pulumi/pulumi/pkg/v3 => ../../pulumi/pulumi/pkg
	github.com/pulumi/pulumi/sdk/v3 => ../../pulumi/pulumi/sdk
)
