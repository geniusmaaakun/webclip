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

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

//バリデーション
//トランザクションの開始
//ログ
//DBへのアクセス
//加工

/*
type markdownUsecase interface {
	Create(title, path, srcUrl string) error
	Delete(title string) error
	DeleteByPath(path string) error
	//DeleteById(id int) error

	//GetByTitle(title string) (*models.MarkdownMemo, error)
	//GetByPath(path string) (*models.MarkdownMemo, error)
	//GetById(id int) (*models.MarkdownMemo, error)
	FindAll() ([]*models.MarkdownMemo, error)
}
*/

type MarkdownUsecase interface {
	Create(title, path, srcUrl string) error
	Delete(title string) error
	DeleteByPath(path string) error
	//DeleteById(id int) error
	FindAll() ([]*models.MarkdownMemo, error)
	FindByTitle(title string) ([]*models.MarkdownMemo, error)
	FindByPath(path string) ([]*models.MarkdownMemo, error)
	FindById(idStr string) (*models.MarkdownMemo, error)
	DeleteIfNotExistsByPath() error
	SearchByTitle(title string) ([]*models.MarkdownMemo, error)
	SearchByBody(bodyStr string) ([]*models.MarkdownMemo, map[string][]string, error)
	CreateZipFile(mds []*models.MarkdownMemo) error
}

type MarkdownInteractor struct {
	txRepo       TransactionRepo
	markdownRepo MarkdownRepo
}

func NewMarkdownInteractor(txRepo TransactionRepo, mdRepo MarkdownRepo) MarkdownUsecase {
	return &MarkdownInteractor{txRepo: txRepo, markdownRepo: mdRepo}
}

//テストしやすい様に、インスタンスを返す様にした方がいいかも
func (u *MarkdownInteractor) Create(title, path, srcUrl string) error {
	md := models.NewMarkdownMemo(title, path, srcUrl)
	//バリデーション
	//https://zenn.dev/mattn/articles/893f28eff96129
	err := validation.ValidateStruct(
		md,
		validation.Field(&md.Title, validation.Required.Error("タイトルは必須入力です"), validation.Length(1, 255)),
		validation.Field(&md.Path, validation.Required.Error("ファイルパスが存在しません"), validation.Length(1, 255)),
		validation.Field(&md.SrcUrl, validation.Required.Error("URLは必須入力です"), is.URL),
	)
	if err != nil {
		return err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return err
	}
	return u.markdownRepo.Create(tx, md)
}

func (u *MarkdownInteractor) Delete(title string) error {
	md := models.NewMarkdownMemo(title, "", "")
	err := validation.ValidateStruct(
		md,
		validation.Field(&md.Title, validation.Required.Error("タイトルは必須入力です"), validation.Length(1, 255)),
	)
	if err != nil {
		return err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return err
	}
	//文字列を渡す
	return u.markdownRepo.DeleteByTitle(tx, md.Title)
}

func (u *MarkdownInteractor) DeleteByPath(path string) error {
	md := models.NewMarkdownMemo("", path, "")
	err := validation.ValidateStruct(
		md,
		validation.Field(&md.Path, validation.Required.Error("ファイルパスが存在しません"), validation.Length(1, 255)),
	)
	if err != nil {
		return err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return err
	}
	return u.markdownRepo.DeleteByPath(tx, md.Path)
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
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, err
	}
	mds, err := u.markdownRepo.FindAll(tx)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindByTitle(title string) ([]*models.MarkdownMemo, error) {
	md := models.NewMarkdownMemo(title, "", "")
	err := validation.ValidateStruct(
		md,
		validation.Field(&md.Title, validation.Required.Error("タイトルは必須入力です"), validation.Length(1, 255)),
	)
	if err != nil {
		return nil, err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, err
	}
	mds, err := u.markdownRepo.FindByTitle(tx, title)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindByPath(path string) ([]*models.MarkdownMemo, error) {
	md := models.NewMarkdownMemo("", path, "")
	err := validation.ValidateStruct(
		md,
		validation.Field(&md.Path, validation.Required.Error("ファイルパスが存在しません"), validation.Length(1, 255)),
	)
	if err != nil {
		return nil, err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, err
	}
	mds, err := u.markdownRepo.FindByPath(tx, path)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) FindById(idStr string) (*models.MarkdownMemo, error) {
	err := validation.Validate(idStr,
		validation.Required, // not empty
		is.Int,
	)
	if err != nil {
		//独自エラーを返す
		return nil, err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//独自エラーを返す
		return nil, err
	}
	err = validation.Validate(id,
		validation.Required, // not empty
		is.Int,
	)
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, err
	}
	md, err := u.markdownRepo.FindById(tx, id)
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
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return err
	}
	mds, err := u.markdownRepo.FindAll(tx)
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
				err := u.markdownRepo.DeleteByPath(tx, md.Path)
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
				err := u.markdownRepo.DeleteByPath(tx, md.Path)
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
	err := validation.Validate(title,
		validation.Required, // not empty
		validation.Length(1, 255),
		is.Alphanumeric,
	)
	if err != nil {
		//独自エラーを返す
		return nil, err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, err
	}
	mds, err := u.markdownRepo.SearchByTitle(tx, title)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func (u *MarkdownInteractor) SearchByBody(bodyStr string) ([]*models.MarkdownMemo, map[string][]string, error) {
	err := validation.Validate(bodyStr,
		validation.Required, // not empty
		validation.Length(1, 255),
		is.Alphanumeric,
	)
	if err != nil {
		//独自エラーを返す
		return nil, nil, err
	}
	tx, err := u.txRepo.NewTransaction(false)
	if err != nil {
		return nil, nil, err
	}
	mds, err := u.markdownRepo.FindAll(tx)
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
func (u *MarkdownInteractor) CreateZipFile(mds []*models.MarkdownMemo) error {
	if len(mds) == 0 || mds == nil {
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
