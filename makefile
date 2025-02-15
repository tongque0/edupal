.PHONY: gen
gen:
	 @cwgo server  --type HTTP --server_name ${svc} --module github.com/tongque0/edupal  --idl ./idl/${svc}.proto
	 @echo "generate ${svc} server success"

# 构建docker镜像命令
# 在使用该命令之前，确保已经登录到github的docker仓库
# docker login ghcr.io -u 用户名 -p <token>
.PHONY: build
build:
	@docker build -f ./deploy/dockerfile.backend -t ghcr.io/tongque0/edupal/edupal:latest .
	@docker push ghcr.io/tongque0/edupal/edupal:latest

.PHONY: build-frontend
build-frontend:
	@cd frontend && pnpm run build
	@docker build -f ./deploy/dockerfile.frontend -t ghcr.io/tongque0/edupal/edupal-frontend:latest ./frontend
	@docker push ghcr.io/tongque0/edupal/edupal-frontend:latest



