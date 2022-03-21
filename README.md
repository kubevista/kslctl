# Kubernetes Serverless CLI (kslctl)

![GitHub](https://img.shields.io/github/license/kubevista/kslctl)

The Kubernetes Serverless CLI Project provides a unifying command-line interface for interacting with the Kubernetes Serverless Stack which constitutes of Knative, Tekton and ArgoCD.

## Installing `kslctl`

### Source install

  If you have [go](https://golang.org/) installed and you want to compile the CLI from source, you can checkout the [Git repository](https://github.com/kubevista/kslctl) and run the following commands:

  ```shell
  go install
  ```

## Useful Commands

The following commands help you understand and effectively use the Kubernetes Serverless CLI:

* `kslctl help:` Displays a list of the commands with helpful information.
* `kslctl install --name=<project-name>`: Installs the Kubernetes Serverless Stack (Knative, Tekton and ArgoCD).
* `kslctl get webhooks:` Displays the webhook URL to be used by the Event Listener Service of Tekton Triggers.
* `kslctl get service-url --name=<service-name> --namespace=<namespace>:` Displays the URL of the Knative Service of the specified name and namespace.
* `kslctl get argocd-password:` Displays the ArgoCD password.





