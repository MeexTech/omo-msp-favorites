.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/favorite.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/repertory.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/activity.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/article.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/notice.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/message.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/display.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/sheet.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/words.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/product.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/favorite/template.proto
