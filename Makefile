.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/favorite.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/repertory.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/activity.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/article.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/notice.proto
