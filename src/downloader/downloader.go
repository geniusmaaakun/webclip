package downloader

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

type ImageDownloader struct {
	Client     *http.Client
	OutputDir  string
	imgCounter int
}

func (d *ImageDownloader) Download(imgURL string) (string, error) {
	/*
		resp, err := d.Client.Get(imgURL)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
		}

		filename := imgFilename(imgURL)
		file, err := os.Create(filepath.Join(d.OutputDir, filename))
		if err != nil {
			return "", err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "", err
		}

		return filename, nil
	*/

	//画像ファイルの拡張子をより安全に取得
	resp, err := d.Client.Get(imgURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")

	//mimeパッケージを使用してContent-Typeから拡張子を取得する方法を以下に示します。
	//mime.ExtensionsByType関数は、Content-Typeに対応する拡張子のスライスを返します。通常は最初の拡張子（ext[0]）を使用しますが、必要に応じて他の拡張子を選択することもできます。
	// これらの方法を使用することで、より正確な画像ファイルの拡張子を取得することができます。ただし、いずれの方法でもサーバーが正しいContent-Typeを返していることが前提となります。サーバーが間違ったContent-Typeを返す場合は、拡張子の判定が正しく行われないことがあります。
	// Determine the file extension from the content type
	ext, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", fmt.Errorf("unsupported content type: %s", contentType)
	}

	//new imageFilename
	//filename := imgFilename(imgURL, ext[0])
	//old imageFIlename
	//filename := imgFilename(imgURL) + ext[0]

	d.imgCounter++
	filename := fmt.Sprintf("%s-%d%s", filepath.Base(d.OutputDir), d.imgCounter, ext[0])

	//画像ファイル名を [ディレクトリ名]-[番号].png の形式に変更
	file, err := os.Create(filepath.Join(d.OutputDir, filename))

	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}
