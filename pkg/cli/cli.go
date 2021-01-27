package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/j4ng5y/sts++/pkg/identity"
	"github.com/spf13/cobra"
)

func Run() {
	var (
		decodeAuthorizationMessage bool

		assumeRoleGenerateCLISkeleton bool
		assumeRoleRoleARN             string
		assumeRoleRoleSessionName     string
		assumeRolePolicy              string
		assumeRoleExternalID          string
		assumeRoleSerialNumber        string
		assumeRoleTokenCode           string
		assumeRoleCLIInputJSON        string
		assumeRolePolicyARNs          []string
		assumeRoleTags                []string
		assumeRoleTransitiveTagKeys   []string
		assumeRoleDuration            time.Duration

		assumeRoleWithSAMLRoleARN             string
		assumeRoleWithSAMLPrincipalARN        string
		assumeRoleWithSAMLSAMLAssertion       string
		assumeRoleWithSAMLPolicy              string
		assumeRoleWithSAMLCLIInputJSON        string
		assumeRoleWithSAMLPolicyARNs          []string
		assumeRoleWithSAMLDuration            time.Duration
		assumeRoleWIthSAMLGenerateCLISkeleton bool

		assumeRoleWithWebIDRoleARN             string
		assumeRoleWithWebIDRoleSessionName     string
		assumeRoleWithWebIDWebIDToken          string
		assumeRoleWithWebIDProviderID          string
		assumeRoleWithWebIDPolicy              string
		assumeRoleWithWebIDCLIInputJSON        string
		assumeRoleWithWebIDPolicyARNs          []string
		assumeRoleWithWebIDDuration            time.Duration
		assumeROleWithWebIDGenerateCLISkeleton bool

		displayFederationToken      bool
		displaySessionToken         bool
		displayAccessKeyInformation bool

		initName   string
		initRegion string
		initOutput string
		initAKID   string
		initSAK    string
	)

	rootCMD := &cobra.Command{
		Use:     "sts",
		Version: "0.1.0",
		Short:   "STS++ is a command-line application that provides a better interface to the AWS STS and related functionality.",
		Run:     func(ccmd *cobra.Command, args []string) {},
	}

	rootCMD.PersistentFlags().BoolVar(&decodeAuthorizationMessage, "decode-authorization-message", false, "Set this flag to decode the authorization message received by AWS.")

	assumeRoleCMD := &cobra.Command{
		Use:   "assume-role",
		Short: "Assume a new AWS IAM role.",
		Run:   func(ccmd *cobra.Command, args []string) {},
	}

	assumeRoleCMD.Flags().StringVarP(&assumeRoleRoleARN, "role-arn", "a", "", "")
	assumeRoleCMD.Flags().StringVarP(&assumeRoleRoleSessionName, "role-session-name", "n", "", "")
	assumeRoleCMD.MarkFlagRequired("role-arn")
	assumeRoleCMD.MarkFlagRequired("role-session-name")
	assumeRoleCMD.Flags().StringSliceVar(&assumeRolePolicyARNs, "policy-arns", []string{}, "")
	assumeRoleCMD.Flags().StringVar(&assumeRolePolicy, "policy", "", "")
	assumeRoleCMD.Flags().DurationVar(&assumeRoleDuration, "duration", 0, "")
	assumeRoleCMD.Flags().StringSliceVar(&assumeRoleTags, "tags", []string{}, "")
	assumeRoleCMD.Flags().StringSliceVar(&assumeRoleTransitiveTagKeys, "transitive-tag-keys", []string{}, "")
	assumeRoleCMD.Flags().StringVar(&assumeRoleExternalID, "external-id", "", "")
	assumeRoleCMD.Flags().StringVar(&assumeRoleSerialNumber, "serial-number", "", "")
	assumeRoleCMD.Flags().StringVar(&assumeRoleTokenCode, "token-code", "", "")
	assumeRoleCMD.Flags().StringVar(&assumeRoleCLIInputJSON, "cli-input-json", "", "")
	assumeRoleCMD.Flags().BoolVar(&assumeRoleGenerateCLISkeleton, "generate-cli-skeleton", false, "")

	withSAMLCMD := &cobra.Command{
		Use:   "with-saml",
		Short: "Assume a new AWS IAM role using a SAML identity.",
		Run:   func(ccmd *cobra.Command, args []string) {},
	}

	withSAMLCMD.Flags().StringVarP(&assumeRoleWithSAMLRoleARN, "role-arn", "a", "", "")
	withSAMLCMD.Flags().StringVarP(&assumeRoleWithSAMLPrincipalARN, "principal-arn", "p", "", "")
	withSAMLCMD.Flags().StringVarP(&assumeRoleWithSAMLSAMLAssertion, "saml-assertion", "s", "", "")
	withSAMLCMD.MarkFlagRequired("role-arn")
	withSAMLCMD.MarkFlagRequired("principal-arn")
	withSAMLCMD.MarkFlagRequired("saml-assertion")
	withSAMLCMD.Flags().StringSliceVar(&assumeRoleWithSAMLPolicyARNs, "policy-arns", []string{}, "")
	withSAMLCMD.Flags().StringVar(&assumeRoleWithSAMLPolicy, "policy", "", "")
	withSAMLCMD.Flags().DurationVar(&assumeRoleWithSAMLDuration, "duration", 0, "")
	withSAMLCMD.Flags().StringVar(&assumeRoleWithSAMLCLIInputJSON, "cli-input-json", "", "")
	withSAMLCMD.Flags().BoolVar(&assumeRoleWIthSAMLGenerateCLISkeleton, "generate-cli-skeleton", false, "")

	withWEBIDCMD := &cobra.Command{
		Use:   "with-web-identity",
		Short: "Assume a new AWS IAM role using a web identity.",
		Run:   func(ccmd *cobra.Command, args []string) {},
	}

	withWEBIDCMD.Flags().StringVarP(&assumeRoleWithWebIDRoleARN, "role-arn", "a", "", "")
	withWEBIDCMD.Flags().StringVarP(&assumeRoleWithWebIDRoleSessionName, "role-session-name", "n", "", "")
	withWEBIDCMD.Flags().StringVarP(&assumeRoleWithWebIDWebIDToken, "web-identity-token", "t", "", "")
	withWEBIDCMD.MarkFlagRequired("role-arn")
	withWEBIDCMD.MarkFlagRequired("role-session-name")
	withWEBIDCMD.MarkFlagRequired("web-identity-token")
	withWEBIDCMD.Flags().StringVar(&assumeRoleWithWebIDProviderID, "provider-id", "", "")
	withWEBIDCMD.Flags().StringSliceVar(&assumeRoleWithWebIDPolicyARNs, "policy-arns", []string{}, "")
	withWEBIDCMD.Flags().StringVar(&assumeRoleWithWebIDPolicy, "policy", "", "")
	withWEBIDCMD.Flags().DurationVar(&assumeRoleWithWebIDDuration, "duration", 0, "")
	withWEBIDCMD.Flags().StringVar(&assumeRoleWithWebIDCLIInputJSON, "cli-input-json", "", "")
	withWEBIDCMD.Flags().BoolVar(&assumeROleWithWebIDGenerateCLISkeleton, "generate-cli-skeleton", false, "")

	displayCMD := &cobra.Command{
		Use:   "display",
		Short: "Display your current AWS credentials.",
		Run:   func(ccmd *cobra.Command, args []string) {},
	}

	displayCMD.Flags().BoolVar(&displayAccessKeyInformation, "access-key-info", false, "Display only access key information.")
	displayCMD.Flags().BoolVar(&displayFederationToken, "federation-token", false, "Display only the federation token.")
	displayCMD.Flags().BoolVar(&displaySessionToken, "session-token", false, "Display only the session token.")

	initCMD := &cobra.Command{
		Use:   "init",
		Short: "Initialize your connection to the AWS backend.",
		Run: func(ccmd *cobra.Command, args []string) {
			if initRegion == "" {
				fmt.Println("Please enter the region your account uses:")
				fmt.Scanln(&initRegion)
			}
			if initOutput == "" {
				fmt.Println("Please enter the output type your account should use:")
				fmt.Scanln(&initOutput)
			}
			if initAKID == "" {
				fmt.Println("Please enter your AWS Access Key ID:")
				fmt.Scanln(&initAKID)
			}
			if initSAK == "" {
				fmt.Println("Please enter your AWS Secret Access Key:")
				fmt.Scanln(&initSAK)
			}

			I, err := identity.New(
				initName,
				identity.Region(initRegion),
				identity.Output(initOutput),
				identity.AWSAccessKeyID(initAKID),
				identity.AWSSecretAccessKey(initSAK))
			if err != nil {
				log.Fatal(err)
			}

			if _, err := I.WriteConfigFile(nil); err != nil {
				log.Fatal(err)
			}
			if _, err := I.WriteCredentialFile(nil); err != nil {
				log.Fatal(err)
			}
		},
	}

	initCMD.Flags().StringVarP(&initName, "name", "n", "default", "")
	initCMD.Flags().StringVarP(&initRegion, "region", "r", "", "")
	initCMD.Flags().StringVarP(&initOutput, "output", "o", "", "")
	initCMD.Flags().StringVarP(&initAKID, "aws-access-key-id", "a", "", "")
	initCMD.Flags().StringVarP(&initSAK, "aws-secret-access-key", "s", "", "")

	rootCMD.AddCommand(assumeRoleCMD, displayCMD, initCMD)
	assumeRoleCMD.AddCommand(withSAMLCMD, withWEBIDCMD)

	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
