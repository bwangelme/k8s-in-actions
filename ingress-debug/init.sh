kubectl cluster-info


if [[ ! $(kubectl get namespace qae) ]]; then
  kubectl create namespace qae
  kubectl create namespace qae-stage
fi

kubectl apply -f ./fruit-deploys/apple.yaml
kubectl apply -f ./fruit-deploys/apple-stage.yaml
kubectl apply -f ./ingress-debug/ing.yaml
