package usecases

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"webclip/src/server/models"
)

//バリデーション
//トランザクションの開始
//ログ
//DBへのアクセス
//加工

type MarkdownUsecase interface {
}

type MarkdownInteractor struct {
	markdownRepo *models.MarkdownRepo
}

func NewMarkdownInteractor(mdRepo *models.MarkdownRepo) *MarkdownInteractor {
	return &MarkdownInteractor{mdRepo}
}

func (mi *MarkdownInteractor) Create(title, path, srcUrl string) error {
	md := models.NewMarkdownMemo(title, path, srcUrl)
	return mi.markdownRepo.Create(md)
}

func (mi *MarkdownInteractor) Delete(title string) error {
	md := models.NewMarkdownMemo(title, "", "")
	return mi.markdownRepo.DeleteByTitle(md)
}

func (mi *MarkdownInteractor) DeleteByPath(path string) error {
	md := models.NewMarkdownMemo("", path, "")
	return mi.markdownRepo.DeleteByPath(md)
}

/*
func (mi *MarkdownInteractor) DeleteById(id int) error {
	md := models.NewMarkdownMemo("", "", "")
	return mi.markdownRepo.Delete(md)
}

func (mi *MarkdownInteractor) GetByTitle(title string) (*models.MarkdownMemo, error) {
	md := models.NewMarkdownMemo(title, "", "")
	return mi.markdownRepo.GetByTitle(md)
}

func (mi *MarkdownInteractor) GetByPath(path string) (*models.MarkdownMemo, error) {
	md := models.NewMarkdownMemo("", path, "")
	return mi.markdownRepo.GetByPath(md)
}

func (mi *MarkdownInteractor) GetById(id int) (*models.MarkdownMemo, error) {
	md := models.NewMarkdownMemo("", "", "")
	return mi.markdownRepo.GetById(md)
}
*/

func (u *MarkdownInteractor) FindAll() ([]*models.MarkdownMemo, error) {
	mds, err := u.markdownRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindByTitle(title string) ([]*models.MarkdownMemo, error) {
	mds, err := u.markdownRepo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindByPath(path string) ([]*models.MarkdownMemo, error) {
	mds, err := u.markdownRepo.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindById(idStr string) (*models.MarkdownMemo, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//独自エラーを返す
		return nil, err
	}
	md, err := u.markdownRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return md, nil
}

/*
func (u *MarkdownInteractor) FindBySrcUrl(srcUrl string) ([]byte, error) {
	md, err := u.markdownRepo.FindBySrcUrl(srcUrl)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(md)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
*/

//cleanコマンドで実行
func (u *MarkdownInteractor) DeleteIfNotExistsByPath() error {
	mds, err := u.markdownRepo.FindAll()
	if err != nil {
		return err
	}
	if mds == nil {
		return nil
	}

	//存在しないか、ディrテクトリの場合削除
	for _, md := range mds {
		file, err := os.Open(md.Path)
		//ファイルが存在しない場合
		if err != nil {
			if os.IsNotExist(err) {
				log.Println(fmt.Sprintf("delete data: %s", md.Path))
				err := u.markdownRepo.DeleteByPath(md)
				if err != nil {
					return err
				}
			}
			//ファイルが存在する場合　ディレクトリの場合
		} else {
			info, err := file.Stat()
			if err != nil {
				return err
			}
			if info.IsDir() {
				log.Println(fmt.Sprintf("delete data: %s", md.Path))
				err := u.markdownRepo.DeleteByPath(md)
				if err != nil {
					return err
				}
			}
		}

		defer file.Close()
	}
	return nil
}

//searchコマンドで実行
func (u *MarkdownInteractor) SearchByTitle(title string) ([]*models.MarkdownMemo, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}
	mds, err := u.markdownRepo.SearchByTitle(title)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (m *MarkdownInteractor) SearchByBody(bodyStr string) ([]*models.MarkdownMemo, map[string][]string, error) {
	if bodyStr == "" {
		return nil, nil, errors.New("body is empty")
	}
	mds, err := m.markdownRepo.FindAll()
	if err != nil {
		return nil, nil, err
	}

	var result []*models.MarkdownMemo
	resultBodyInFile := make(map[string][]string)

	for _, md := range mds {
		file, err := os.Open(md.Path)
		if err != nil {
			return nil, nil, err
		}
		defer file.Close()
		body, err := io.ReadAll(file)
		if err != nil {
			return nil, nil, err
		}
		//含む場合
		if strings.Contains(string(body), bodyStr) {
			result = append(result, md)
			//ファイル内の該当行を取得
			i := 0
			b := bytes.NewBuffer(body)
			sc := bufio.NewScanner(b)
			for sc.Scan() {
				if strings.Contains(sc.Text(), bodyStr) {
					resultBodyInFile[md.Path] = append(resultBodyInFile[md.Path], fmt.Sprintf("line:%d %s", i, sc.Text()))
				}
				i++
			}
		}
	}
	return result, resultBodyInFile, nil
}

//配列のファイルリストをzipに圧縮する
func createZip(files []string, zipFilename string) error {
	//zipファイルを作成
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	//zip.Writerを作成
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	//ファイルをzipに追加
	for _, file := range files {
		if err := addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	return nil
}

//zipにファイルを追加する
func addFileToZip(zipWriter *zip.Writer, file string) error {
	//ファイルを開く
	srcFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	//ファイル情報を取得
	fileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	//ヘッダーを作成 ファイル情報を渡す
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	//ヘッダーのNameを設定
	header.Name = filepath.Join(filepath.Base(filepath.Dir(file)), filepath.Base(file))
	//ヘッダーのMethodを設定
	header.Method = zip.Deflate

	//ヘッダーを元にファイルをzipに書き込む
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	//ファイルをzipに書き込む。ここに書き込むとzipファイルに書き込まれる
	_, err = io.Copy(writer, srcFile)
	return err
}

//ファイルを一箇所に集めるためにzip化
func (m *MarkdownInteractor) CreateZipFile(mds []*models.MarkdownMemo) error {
	if len(mds) == 0 {
		return errors.New("no data")
	}

	var files []string
	for _, md := range mds {
		files = append(files, md.Path)
	}

	//zip化
	err := createZip(files, "webclip.zip")
	if err != nil {
		return err
	}
	return nil
}
