package product_adm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/entity"
	"github.com/private-project-pp/product-rpc-service/mocks/postgre_repo_mocks"
	"github.com/private-project-pp/product-rpc-service/shared/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDbMock() (*gorm.DB, *sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, nil, err
	}

	return gormDB, db, mock, nil
}

func TestAddingNewProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dbConn, db, dbMock, _ := setupDbMock()
	defer db.Close()
	type args struct {
		productRepo *postgre_repo_mocks.MockProduct
		db          *gorm.DB
		dbMock      sqlmock.Sqlmock
	}
	test_case := []struct {
		name    string
		args    args
		usecase func(args)
		wants   func(*testing.T, *model.AddProductResponse, error)
		payload *model.AddProductRequest
	}{
		{
			name: "Success add new product",
			args: args{
				productRepo: postgre_repo_mocks.NewMockProduct(mockCtrl),
				db:          dbConn,
				dbMock:      dbMock,
			},
			usecase: func(a args) {
				a.dbMock.ExpectBegin()
				a.productRepo.EXPECT().AddProduct(gomock.Any(), gomock.AssignableToTypeOf(entity.Product{})).Return(nil).Times(1)
				a.productRepo.EXPECT().BulkAddProductImages(gomock.Any(), gomock.AssignableToTypeOf([]entity.ProductsImages{})).Return(nil).Times(1)
				a.dbMock.ExpectCommit()

			},
			wants: func(t *testing.T, apr *model.AddProductResponse, err error) {
				assert.NotEqual(t, apr.GetSecureId(), "", "SecureId shouldly exists")
				assert.Equal(t, err, nil, "Err shouldly NIL")
			},
			payload: &model.AddProductRequest{
				Name:    "Produk test",
				Barcode: "019283918282",
				Status:  0,
				Note:    "Catatan",
				Unit:    utils.GenerateUUID(),
				Image:   utils.GenerateUUID(),
				Images: []*model.UploadedProductImageRequest{
					{
						SecureId: utils.GenerateUUID(),
						FileName: "tes-tambah-produk",
						FileType: ".png",
					},
				},
			},
		},
		{
			name: "Failed add new product",
			args: args{
				productRepo: postgre_repo_mocks.NewMockProduct(mockCtrl),
				db:          dbConn,
				dbMock:      dbMock,
			},
			usecase: func(a args) {
				err := errors.New("Ada error")
				err = stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
				a.dbMock.ExpectBegin()
				a.productRepo.EXPECT().AddProduct(gomock.Any(), gomock.AssignableToTypeOf(entity.Product{})).Return(err).Times(1)
				a.dbMock.ExpectRollback()

			},
			wants: func(t *testing.T, apr *model.AddProductResponse, err error) {
				assert.Equal(t, apr.GetSecureId(), "", "SecureId shouldly empty")
				assert.NotEqual(t, err, nil, "Err shouldly not NIL")
			},
			payload: &model.AddProductRequest{
				Name:    "Produk test",
				Barcode: "019283918282",
				Status:  0,
				Note:    "Catatan",
				Unit:    utils.GenerateUUID(),
				Image:   utils.GenerateUUID(),
				Images: []*model.UploadedProductImageRequest{
					{
						SecureId: utils.GenerateUUID(),
						FileName: "tes-tambah-produk",
						FileType: ".png",
					},
				},
			},
		},
		{
			name: "Failed add new product [IMAGES]",
			args: args{
				productRepo: postgre_repo_mocks.NewMockProduct(mockCtrl),
				db:          dbConn,
				dbMock:      dbMock,
			},
			usecase: func(a args) {
				err := errors.New("Ada error")
				err = stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
				a.dbMock.ExpectBegin()
				a.productRepo.EXPECT().AddProduct(gomock.Any(), gomock.AssignableToTypeOf(entity.Product{})).Return(nil).Times(1)
				a.productRepo.EXPECT().BulkAddProductImages(gomock.Any(), gomock.AssignableToTypeOf([]entity.ProductsImages{})).Return(err).Times(1)
				a.dbMock.ExpectRollback()

			},
			wants: func(t *testing.T, apr *model.AddProductResponse, err error) {
				assert.Equal(t, apr.GetSecureId(), "", "SecureId shouldly empty")
				assert.NotEqual(t, err, nil, "Err shouldly not NIL")
			},
			payload: &model.AddProductRequest{
				Name:    "Produk test",
				Barcode: "019283918282",
				Status:  0,
				Note:    "Catatan",
				Unit:    utils.GenerateUUID(),
				Image:   utils.GenerateUUID(),
				Images: []*model.UploadedProductImageRequest{
					{
						SecureId: utils.GenerateUUID(),
						FileName: "tes-tambah-produk",
						FileType: ".png",
					},
				},
			},
		},
		{
			name: "Failed COMMIT add new product",
			args: args{
				productRepo: postgre_repo_mocks.NewMockProduct(mockCtrl),
				db:          dbConn,
				dbMock:      dbMock,
			},
			usecase: func(a args) {
				err := errors.New("Ada error")
				err = stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
				a.dbMock.ExpectBegin()
				a.productRepo.EXPECT().AddProduct(gomock.Any(), gomock.AssignableToTypeOf(entity.Product{})).Return(nil).Times(1)
				a.productRepo.EXPECT().BulkAddProductImages(gomock.Any(), gomock.AssignableToTypeOf([]entity.ProductsImages{})).Return(nil).Times(1)
				a.dbMock.ExpectCommit().WillReturnError(err)
				a.dbMock.ExpectRollback()

			},
			wants: func(t *testing.T, apr *model.AddProductResponse, err error) {
				assert.Equal(t, apr.GetSecureId(), "", "SecureId shouldly empty")
				assert.NotEqual(t, err, nil, "Err shouldly not NIL")
			},
			payload: &model.AddProductRequest{
				Name:    "Produk test",
				Barcode: "019283918282",
				Status:  0,
				Note:    "Catatan",
				Unit:    utils.GenerateUUID(),
				Image:   utils.GenerateUUID(),
				Images: []*model.UploadedProductImageRequest{
					{
						SecureId: utils.GenerateUUID(),
						FileName: "tes-tambah-produk",
						FileType: ".png",
					},
				},
			},
		},
	}

	for _, test := range test_case {
		_ = test
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
			testedObj := productAdm{
				productRepo: test.args.productRepo,
				dbTx:        test.args.db,
			}
			test.usecase(test.args)
			result, err := testedObj.AddingNewProduct(context.Background(), test.payload)
			test.wants(t, result, err)
		})
	}
}
