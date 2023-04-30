NAME = webclip

# ディレクトリを作成し、DBを作成する。ルートに作成
install: 

test: 
	go test ./src/actions -v
	go test ./src/wcconverter -v
	go test ./src/wcdownloader -v
	go test ./src/server/models/rdb -v
	go test ./src/server/usecases -v
	go test ./src/server/controllers/handler -v

.PHONY: install test