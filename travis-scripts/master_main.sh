#!/usr/bin/env bash

# Delete K8s cluster
gcloud container clusters delete $CLUSTER_NAME_NIGHTLY --zone $CLOUDSDK_COMPUTE_ZONE --project $PROJECT_NAME --quiet

source ./travis-scripts/setup_functions.sh

# prints the full command before output of the command.
set -x

install_hub
install_yq
setup_gcloud
setup_gcloud_master
install_sed

# TODO: update project name(s)
hub delete -y $GITHUB_ORG_NIGHTLY/carts || true

# TODO: For developing purposes the branch 'travis-nightly-build' is selected. Afterwards, change it back to master.
git clone --branch travis-nightly-build https://github.com/keptn/keptn
cd keptn/install/scripts

source ./defineCredentials.sh

# Set enviornment variables used in replaceCreds function
GITU = $GITHUB_USER_NAME_NIGHTLY
GITAT = $GITHUB_TOKEN_NIGHTLY
GITE = $GITHUB_EMAIL_NIGHTLY
CLN = $CLUSTER_NAME_NIGHTLY
CLZ = $CLOUDSDK_COMPUTE_ZONE
PROJ = $PROJECT_NAME
GITO = $GITHUB_ORG_NIGHTLY

replaceCreds

./installKeptn.sh
cd ../..

# Test front-end keptn v.0.1
# export FRONT_END_DEV=$(kubectl describe svc front-end -n dev | grep "LoadBalancer Ingress:" | sed 's~LoadBalancer Ingress:[ \t]*~~')
# export FRONT_END_STAGING=$(kubectl describe svc front-end -n staging | grep "LoadBalancer Ingress:" | sed 's~LoadBalancer Ingress:[ \t]*~~')
# export FRONT_END_PRODUCTION=$(kubectl describe svc front-end -n production | grep "LoadBalancer Ingress:" | sed 's~LoadBalancer Ingress:[ \t]*~~')

export ISTIO_INGRESS=$(kubectl describe svc istio-ingressgateway -n istio-system | grep "LoadBalancer Ingress:" | sed 's~LoadBalancer Ingress:[ \t]*~~')
export_names

#- cat ../test/keptn.postman_environment.json |sed 's~FRONT_END_DEV_PLACEHOLDER~'"$FRONT_END_DEV"'~' |sed 's~FRONT_END_STAGING_PLACEHOLDER~'"$FRONT_END_STAGING"'~' |sed 's~FRONT_END_PRODUCTION_PLACEHOLDER~'"$FRONT_END_PRODUCTION"'~' |sed 's~ISTIO_INGRESS_PLACEHOLDER~'"$ISTIO_INGRESS"'~' >> ../test/env.json
#- npm install newman
#- node_modules/.bin/newman run ../test/keptn.postman_collection.json -e ../test/env.json
        
# Clean up cluster
#- ./cleanupCluster.sh

