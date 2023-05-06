package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewImportExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *ImportExcelLogic {
	return &ImportExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *ImportExcelLogic) ImportExcel(req *types.ImportExcelReq) (resp *types.ImportExcelResp, err error) {
	file, _, err := l.r.FormFile("file")
	if err != nil {
		response.Error(450, err.Error())
	}

	// 读取excel流
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		response.Error(450, "open excel error:"+err.Error())
	}
	// 解析excel数据
	usersInfo, lxRrr := readExcel(xlsx)
	if lxRrr != nil {
		response.Error(450, "解析excel数据失败:"+err.Error())
	}


	var result []types.User // 将models.Users转换成types.User
	if len(usersInfo) > 0 {
		for _, usersmodel := range usersInfo {
			var typeUser types.User
			_ = copier.Copy(&typeUser, usersmodel)
			result = append(result, typeUser)
		}
	}
	return &types.ImportExcelResp{
		List: result,
	}, nil
}

// readExcel：读取excel转成切片
func readExcel(xlsx *excelize.File) ([]models.Users, error) {
	// 根据名字获取cells的内容，返回的是一个[][]string
	// GetRows按给定的工作表名称返回工作表中的所有行，返回为二维数组
	var lxProduct []models.Users
	for _, sheetName := range xlsx.GetSheetList() {
		rows, err := xlsx.GetRows(sheetName)
		if err != nil {
			response.Error(450, "excel get rows error:"+err.Error())
		}

		for i, row := range rows {
			// 去掉第一行是excel表头部分
			if i == 0 {
				continue
			}
			var data models.Users
			for k, v := range row {
				// 第一列是电话
				if k == 0 {
					data.Mobile = v
				}
				// 第二列是姓名
				if k == 1 {
					data.Username = v
				}
				if k == 2 {
					data.Password = v
				}
				if k == 3 {
					data.Sex, err = strconv.ParseInt(v, 10, 64)
				}
				if k == 4 {
					data.Email = v
				}
				if k == 5 {
					data.Address = v
				}
				if k == 6 {
					data.Birthday = v
				}
			}
			// 添加到切片中
			lxProduct = append(lxProduct, data)
		}
	}

	// 声明一个数组
	return lxProduct, nil
}
