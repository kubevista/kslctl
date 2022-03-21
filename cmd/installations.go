package cmd

var knativeInstallationScript = `
#!/bin/bash

echo "Installing Knative Serving CRDs..."
# kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-crds.yaml
# echo "Finished Installing Knative Serving CRDs"
# echo "Installing core components of Knative Serving..."
# kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-core.yaml
# echo "Finished Installing core components of Knative Serving"
# echo "Installing networking layer (Knative Kourier controller)..."
# kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.3.0/kourier.yaml
# echo "Finished Installing networking layer"
# echo "Patching networking layer settings..."
# kubectl patch configmap/config-network \
#   --namespace knative-serving \
#   --type merge \
#   --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
# echo "Installing Magic DNS (sslip.io)..."
# kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-default-domain.yaml
echo "Finished Installing Magic DNS"
`

var tektonInstallationScript = `
#!/bin/bash

echo "Installing Tekton Pipelines..."
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.26.0/release.yaml
# echo "Finished Installing Tekton Pipelines"
# echo "Installing Tekton Dashboard..."
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml
echo "Finished Installing Tekton Dashboard"
# echo "Installing Tekton Triggers..."
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/release.yaml
# echo "Finished Installing Tekton Triggers"
# echo "Installing Tekton Core Interceptors..."
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml
# echo "Finished Installing Tekton Core Interceptors"
# echo "Installing Nginx-Ingress Controller..."
# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/aws/deploy.yaml
# echo "Finished Installing Nginx-Ingress Controller"
`

var argocdInstallationScript = `
#!/bin/bash

echo "Installing ArgoCD..."
# kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
# echo "Finished Installing ArgoCD"
`
