.PHONY: gen
gen:
	 @cwgo server  --type HTTP --server_name ${svc} --module github.com/tongque0/edupal  --idl ./idl/${svc}.proto
	 @echo "generate ${svc} server success"

# .PHONY: build
# build:
# 	@docker build -t tsrms:latest .
# 	@docker tag tsrms:latest serverless-100026543835-docker.pkg.coding.net/trsms/tsrms/tsrms:latest
# 	@docker push serverless-100026543835-docker.pkg.coding.net/trsms/tsrms/tsrms:latest

