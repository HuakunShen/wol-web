buildx:
	docker buildx build --push \
		--platform linux/arm64,linux/amd64 \
		-t huakunshen/wol:latest .

