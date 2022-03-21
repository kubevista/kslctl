/*
Copyright © 2022 Avik <akundu@redhat.com>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Helps to install serverless project on kubernetes",
	Long: `For example:
	kslctl install --name=argocd
	kslctl install --name=knative
	kslctl install --name=tekton
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		defaultNamespace := make(map[string]string)
		defaultNamespace["tekton"] = "tekton-pipelines"
		defaultNamespace["knative"] = "knative-serving"
		defaultNamespace["argocd"] = "argocd"

		if kubectlPresentCheck() {
			fmt.Println("kubectl is installed.")
			fmt.Println("Installing Project:", name, "in namespace:", defaultNamespace[name])
			installProject(name, defaultNamespace[name])
		} else {
			fmt.Println("kubectl is not installed. Please install kubectl and try again.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("name", "n", "", "Name of project to install")

	installCmd.MarkFlagRequired("name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func kubectlPresentCheck() bool {
	topCommand := "kubectl"
	arg0 := "version"

	cmd := exec.Command(topCommand, arg0)
	_, err := cmd.Output()
	return err == nil
}

func installProject(name string, namespace string) {
	switch name {
	case "argocd":
		installArgoCD(namespace)
	case "knative":
		installKnative(namespace)
	case "tekton":
		installTektonPipelines(namespace)
	case "all":
		installArgoCD(namespace)
		installKnative(namespace)
		installTektonPipelines(namespace)
	default:
		fmt.Println("Tool not found.")
	}
}

func installArgoCD(namespace string) {
	fmt.Println("Installing ArgoCD Started")
	fmt.Println("Creating namespace:", namespace)
	createNamespace(namespace)
	fmt.Println("Namespace created.")
	c := exec.Command("bash")
	c.Stdin = strings.NewReader(argocdInstallationScript)
	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	fmt.Println("ArgoCD installation complete.")
}

func installKnative(namespace string) {
	fmt.Println("Installing Knative Serving Started...")
	c := exec.Command("bash")
	c.Stdin = strings.NewReader(knativeInstallationScript)
	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	fmt.Println("Knative installation complete.")
}

func installTektonPipelines(namespace string) {
	fmt.Println("Installing TektonPipelines Started...")
	c := exec.Command("bash")
	c.Stdin = strings.NewReader(tektonInstallationScript)
	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	fmt.Println("TektonPipelines installation complete.")
}

func createNamespace(namespace string) {
	topCommand := "kubectl"
	arg0 := "create"
	arg1 := "namespace"
	arg2 := namespace
	// temp := "version"

	cmd := exec.Command(topCommand, arg0, arg1, arg2)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
