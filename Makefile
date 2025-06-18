# Build production image
build:
	docker build -t payment:latest -f deployments/docker/Dockerfile .

# Build development image
build-dev:
	docker build -t payment-dev:latest -f deployments/docker/Dockerfile.dev .

# Run docker-compose with prod profile
run-compose:
	docker-compose -f deployments/docker/docker-compose.yml --profile prod up --build

# Run docker-compose with dev profile
run-compose-dev:
	docker-compose -f deployments/docker/docker-compose.yml --profile dev up

# Push production image to your registry
push:
	docker tag payment:latest your-registry/payment:latest
	docker push your-registry/payment:latest

# Deploy to Kubernetes
deploy-k8s:
	kubectl apply -f deployments/azure/kubernetes/

# Stop and clean up docker-compose (volumes included)
clean-compose:
	docker-compose -f deployments/docker/docker-compose.yml down -v

# Stop and clean up docker-compose for dev profile
clean-compose-dev:
	docker-compose -f deployments/docker/docker-compose.yml --profile dev down -v
