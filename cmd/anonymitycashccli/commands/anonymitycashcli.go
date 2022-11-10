package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/anonymitycash/anonymitycash/util"
)

// anonymitycashcli usage template
var usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletDisable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

  available with wallet enable:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletEnable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// commandError is an error used to signal different error situations in command handling.
type commandError struct {
	s         string
	userError bool
}

func (c commandError) Error() string {
	return c.s
}

func (c commandError) isUserError() bool {
	return c.userError
}

func newUserError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: true}
}

func newSystemError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: false}
}

func newSystemErrorF(format string, a ...interface{}) commandError {
	return commandError{s: fmt.Sprintf(format, a...), userError: false}
}

// Catch some of the obvious user errors from Cobra.
// We don't want to show the usage message for every error.
// The below may be to generic. Time will show.
var userErrorRegexp = regexp.MustCompile("argument|flag|shorthand")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}

// AnonymitycashcliCmd is Anonymitycashcli's root command.
// Every other command attached to AnonymitycashcliCmd is a child command to it.
var AnonymitycashcliCmd = &cobra.Command{
	Use:   "anonymitycashcli",
	Short: "Anonymitycashcli is a commond line client for anonymitycash core (a.k.a. anonymitycashd)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.SetUsageTemplate(usageTemplate)
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command AnonymitycashcliCmd and sets flags appropriately.
func Execute() {

	AddCommands()
	AddTemplateFunc()

	if _, err := AnonymitycashcliCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}

// AddCommands adds child commands to the root command AnonymitycashcliCmd.
func AddCommands() {
	AnonymitycashcliCmd.AddCommand(createAccessTokenCmd)
	AnonymitycashcliCmd.AddCommand(listAccessTokenCmd)
	AnonymitycashcliCmd.AddCommand(deleteAccessTokenCmd)
	AnonymitycashcliCmd.AddCommand(checkAccessTokenCmd)

	AnonymitycashcliCmd.AddCommand(createAccountCmd)
	AnonymitycashcliCmd.AddCommand(deleteAccountCmd)
	AnonymitycashcliCmd.AddCommand(listAccountsCmd)
	AnonymitycashcliCmd.AddCommand(updateAccountAliasCmd)
	AnonymitycashcliCmd.AddCommand(createAccountReceiverCmd)
	AnonymitycashcliCmd.AddCommand(listAddressesCmd)
	AnonymitycashcliCmd.AddCommand(validateAddressCmd)
	AnonymitycashcliCmd.AddCommand(listPubKeysCmd)

	AnonymitycashcliCmd.AddCommand(createAssetCmd)
	AnonymitycashcliCmd.AddCommand(getAssetCmd)
	AnonymitycashcliCmd.AddCommand(listAssetsCmd)
	AnonymitycashcliCmd.AddCommand(updateAssetAliasCmd)

	AnonymitycashcliCmd.AddCommand(getTransactionCmd)
	AnonymitycashcliCmd.AddCommand(listTransactionsCmd)

	AnonymitycashcliCmd.AddCommand(getUnconfirmedTransactionCmd)
	AnonymitycashcliCmd.AddCommand(listUnconfirmedTransactionsCmd)
	AnonymitycashcliCmd.AddCommand(decodeRawTransactionCmd)

	AnonymitycashcliCmd.AddCommand(listUnspentOutputsCmd)
	AnonymitycashcliCmd.AddCommand(listBalancesCmd)

	AnonymitycashcliCmd.AddCommand(rescanWalletCmd)
	AnonymitycashcliCmd.AddCommand(walletInfoCmd)

	AnonymitycashcliCmd.AddCommand(buildTransactionCmd)
	AnonymitycashcliCmd.AddCommand(signTransactionCmd)
	AnonymitycashcliCmd.AddCommand(submitTransactionCmd)
	AnonymitycashcliCmd.AddCommand(estimateTransactionGasCmd)

	AnonymitycashcliCmd.AddCommand(getBlockCountCmd)
	AnonymitycashcliCmd.AddCommand(getBlockHashCmd)
	AnonymitycashcliCmd.AddCommand(getBlockCmd)
	AnonymitycashcliCmd.AddCommand(getBlockHeaderCmd)
	AnonymitycashcliCmd.AddCommand(getDifficultyCmd)
	AnonymitycashcliCmd.AddCommand(getHashRateCmd)

	AnonymitycashcliCmd.AddCommand(createKeyCmd)
	AnonymitycashcliCmd.AddCommand(deleteKeyCmd)
	AnonymitycashcliCmd.AddCommand(listKeysCmd)
	AnonymitycashcliCmd.AddCommand(updateKeyAliasCmd)
	AnonymitycashcliCmd.AddCommand(resetKeyPwdCmd)
	AnonymitycashcliCmd.AddCommand(checkKeyPwdCmd)

	AnonymitycashcliCmd.AddCommand(signMsgCmd)
	AnonymitycashcliCmd.AddCommand(verifyMsgCmd)
	AnonymitycashcliCmd.AddCommand(decodeProgCmd)

	AnonymitycashcliCmd.AddCommand(createTransactionFeedCmd)
	AnonymitycashcliCmd.AddCommand(listTransactionFeedsCmd)
	AnonymitycashcliCmd.AddCommand(deleteTransactionFeedCmd)
	AnonymitycashcliCmd.AddCommand(getTransactionFeedCmd)
	AnonymitycashcliCmd.AddCommand(updateTransactionFeedCmd)

	AnonymitycashcliCmd.AddCommand(isMiningCmd)
	AnonymitycashcliCmd.AddCommand(setMiningCmd)

	AnonymitycashcliCmd.AddCommand(netInfoCmd)
	AnonymitycashcliCmd.AddCommand(gasRateCmd)

	AnonymitycashcliCmd.AddCommand(versionCmd)
}

// AddTemplateFunc adds usage template to the root command AnonymitycashcliCmd.
func AddTemplateFunc() {
	walletEnableCmd := []string{
		createAccountCmd.Name(),
		listAccountsCmd.Name(),
		deleteAccountCmd.Name(),
		updateAccountAliasCmd.Name(),
		createAccountReceiverCmd.Name(),
		listAddressesCmd.Name(),
		validateAddressCmd.Name(),
		listPubKeysCmd.Name(),

		createAssetCmd.Name(),
		getAssetCmd.Name(),
		listAssetsCmd.Name(),
		updateAssetAliasCmd.Name(),

		createKeyCmd.Name(),
		deleteKeyCmd.Name(),
		listKeysCmd.Name(),
		resetKeyPwdCmd.Name(),
		checkKeyPwdCmd.Name(),
		signMsgCmd.Name(),

		buildTransactionCmd.Name(),
		signTransactionCmd.Name(),

		getTransactionCmd.Name(),
		listTransactionsCmd.Name(),
		listUnspentOutputsCmd.Name(),
		listBalancesCmd.Name(),

		rescanWalletCmd.Name(),
		walletInfoCmd.Name(),
	}

	cobra.AddTemplateFunc("WalletEnable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return true
			}
		}
		return false
	})

	cobra.AddTemplateFunc("WalletDisable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return false
			}
		}
		return true
	})
}
