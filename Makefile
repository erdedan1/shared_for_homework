PROTOC = protoc

.PHONY: proto
proto: order spot

order:
	$(PROTOC) -I=proto/order_service/proto \
		--go_out=. \
		--go-grpc_out=. \
		order.proto

spot:
	$(PROTOC) -I=proto/spot_instrument_service/proto \
		--go_out=. \
		--go-grpc_out=. \
		spot_instrument.proto

