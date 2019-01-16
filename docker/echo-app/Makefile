IMAGE_NAME = mmaeda/k8s-echo-app
APP_VERSION = v1

dockerBuild: 
	docker build -t $(IMAGE_NAME):$(APP_VERSION) --build-arg APP_VERSION="$(APP_VERSION)" .
dockerPush: dockerBuild
	docker push $(IMAGE_NAME):$(APP_VERSION)

