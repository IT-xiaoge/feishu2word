package cmd

import (
    "errors"
    "feishu2word/module"
    "github.com/spf13/cobra"
)

var (
    savepath      string
    usertoken   string
)

var rootCmd = &cobra.Command{
    Use:         "feishu2word",
    Short:       "feishu wiki download",
    Run:         downloadALL,
    PreRunE:     checkdownloadallParam,
    Example:     `feishu2word -p xxxx -t xxxx`,
}

func downloadALL(cmd *cobra.Command, args []string) {
    module.GetDoc("",usertoken,savepath)
    module.GetSpaces(usertoken,savepath)
}

func checkdownloadallParam(cmd *cobra.Command, args []string) error {
    if  savepath == "" || usertoken == "" {
        return errors.New("savepath and usertoken are required")
    }
    return nil
}

func Execute() {
    rootCmd.Execute()
}

func init() {
    rootCmd.Flags().StringVarP(&savepath, "savepath", "p", "", "feishu wiki save to path")
    rootCmd.Flags().StringVarP(&usertoken, "usertoken", "t", "", "feishu app user_access_token")
}