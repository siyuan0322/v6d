REGISTRY 			:= registry-vpc.cn-hongkong.aliyuncs.com
IMAGE 				:= libvineyard/vineyardd

ubuntu: ubuntu-build ubuntu-push

ubuntu-build: Dockerfile.ubuntu
	docker build .. -f Dockerfile.ubuntu -t $(REGISTRY)/$(IMAGE):ubuntu
.PHONY: ubuntu-build

ubuntu-push: Dockerfile.ubuntu
	docker push $(REGISTRY)/$(IMAGE):ubuntu
.PHONY: ubuntu-push
