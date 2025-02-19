#!/bin/bash

set -e  # Exit on error

# Variables
IMAGE_NAME="shoedesk"
TAG="latest"
KIND_CLUSTER_NAME="kind"
DEPLOYMENT_YAML="deployment.yaml"

# Step 1: Build Docker Image
echo "Building Docker image..."
docker build -t ${IMAGE_NAME}:${TAG} .

# Step 2: Create a Kind Cluster (if not exists)
if ! kind get clusters | grep -q ${KIND_CLUSTER_NAME}; then
    echo "Creating Kind cluster..."
    kind create cluster #--name ${KIND_CLUSTER_NAME}
else
    echo "Kind cluster already exists."
fi

# Step 3: Load Image into Kind
#echo "Loading image into Kind cluster..."
#kind load docker-image ${IMAGE_NAME}:${TAG} --name ${KIND_CLUSTER_NAME}
# docker save shoedesk:latest -o shoedesk.tar && \
#     kind load image-archive shoedesk.tar
docker tag shoedesk:latest danwroy/shoedesk:latest
# docker push danwroy/shoedesk:latest
    
# Step 4: Apply Kubernetes Deployment
echo "Applying Kubernetes deployment..."
kubectl apply -f ${DEPLOYMENT_YAML}

# Step 5: Verify Deployment
echo "Verifying deployment..."
kubectl get pods -o wide
kubectl get deployments

echo "Deployment completed successfully!"
