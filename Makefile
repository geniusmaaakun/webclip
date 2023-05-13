NAME = webclip

# ディレクトリを作成し、DBを作成する。ルートに作成
install: 
	cd frontend && npm run build
	go install


# OSによって配置場所を変える。手動で変更するマニュアルをREADMEに記載する
#install: build
#	cp $(NAME) /usr/local/bin/$(NAME)

build:
	go build -o $(NAME) main.go

#クロスコンパイル
#https://qiita.com/Utr/items/9469c1611abe8a0a3486
#env GOOS=linux GOARCH=amd64 go build main.go（buildするファイル名）

test: 
	go test ./src/actions -v
	go test ./src/wcconverter -v
	go test ./src/wcdownloader -v
	go test ./src/server/models/rdb -v
	go test ./src/server/usecases -v
	go test ./src/server/controllers/handler -v
	cd frontend && npm test

.PHONY: install test build