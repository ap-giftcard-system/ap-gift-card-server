include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
AP_GIFT_SERVER=ap-gift-card-server
AP_GIFT_PORT=41125
AP_GIFT_TAG=1.0
AP_GIFT_IMAGE=$(DOCKER_USERNAME)/$(AP_GIFT_SERVER):$(AP_GIFT_TAG)
DOCKER_USERNAME=apnailart
DOCKER_RMI=docker rmi -f
DOCKER_RM=docker rm -f
DOCKER_PULL=docker pull
DOCKER_PUSH=docker push
DOCKER_IMAGE_LIST_ID=docker images -q
DOCKER_CONTAINER_LIST_ID=docker ps -aq
DOCKER_BUILD_SCRIPT=docker build --no-cache -t $(AP_GIFT_IMAGE) .
DOCKER_RUN_SCRIPT=docker run -d --rm $\
					--name $(AP_GIFT_SERVER) $\
					--env-file .env $\
					-e GIN_MODE=release $\
					-p $(AP_GIFT_PORT):$(PRODUCTION_PORT) $\
					$(AP_GIFT_IMAGE)

ENV_VARS=$(shell grep -v '^#' .env | xargs)
GOOGLE_CLOUD_PROJECT_ID=ap-nail-art
GOOGLE_CLOUD_REPOSITORY=ap-gift-server
GOOGLE_CLOUD_REGION=us-central1

GOOGLE_CLOUD_BUILD_SCRIPT=gcloud builds submit --tag us-central1-docker.pkg.dev/$(GOOGLE_CLOUD_PROJECT_ID)/$(GOOGLE_CLOUD_REPOSITORY)/$(AP_GIFT_SERVER):$(AP_GIFT_TAG)


GOOGLE_CLOUD_DEPLOY_SCRIPT=gcloud run deploy $(AP_GIFT_SERVER) --source . $\ 
						--image gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(AP_GIFT_SERVER) $\
						--platform managed $\
						--region $(GOOGLE_CLOUD_REGION) $\
						--allow-unauthenticated $\
						--set-env-vars "GIN_MODE=release,$(ENV_VARS)"

#############################################
############### LOCAL BUILD #################
#############################################

# dev-mode
.phony: dev
dev: 
	@$(COMPILEDAEMON_PATH) -command="./$(AP_GIFT_SERVER)"

# local run
.phony: go-run
go-run:
	@go run .


#############################################
############### DOCKER BUILD ################
#############################################
docker-remove-ap-gift-img:
	$(DOCKER_RMI) $(AP_GIFT_IMAGE)

docker-build-ap-gift: docker-remove-ap-gift-img
	$(DOCKER_BUILD_SCRIPT)

docker-run-ap-gift:
	$(DOCKER_RUN_SCRIPT)

docker-update-remote-image: docker-build-ap-gift
	$(DOCKER_PUSH) $(AP_GIFT_IMAGE)

docker-pull-ap-gift:
	$(DOCKER_PULL) $(AP_GIFT_IMAGE)

docker-dev-ap-gift: docker-pull-ap-gift docker-run-ap-gift


.PHONY: docker-clean
docker-clean:
	$(DOCKER_RM) $(AP_GIFT_SERVER) $&
	$(DOCKER_RMI) $(AP_GIFT_IMAGE)

#############################################
######### Google Cloud BUILD ################
#############################################
gcloud-build:
	$(GOOGLE_CLOUD_BUILD_SCRIPT)

gcloud-deploy: gcloud-build
	$(GOOGLE_CLOUD_DEPLOY_SCRIPT)


update-remote-image-docker-gcloud: docker-update-remote-image gcloud-deploy