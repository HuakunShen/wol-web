IMAGE_NAME ?= huakunshen/wol:latest
.phony: buildx buildx-podman
buildx:
	docker buildx build --push \
		--platform linux/arm64,linux/amd64 \
		-t $(IMAGE_NAME) .

buildx-podman:
	podman buildx build --jobs=2 --platform=linux/arm64,linux/amd64 --manifest=$(IMAGE_NAME) .;  \
	podman manifest push $(IMAGE_NAME)
