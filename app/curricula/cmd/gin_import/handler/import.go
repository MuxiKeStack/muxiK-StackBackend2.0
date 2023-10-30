package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/model/import_curricular"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
)

func ImportCurricula(c *gin.Context) { // 从表单中获取上传的文件
	file, err := c.FormFile("curricular")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开上传的文件
	xlFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer xlFile.Close()

	// 读取xlsx文件
	xl, err := xlsx.OpenReaderAt(xlFile, file.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 处理xlsx文件数据
	err = import_curricular.ImportCurricularByXlsx(xl)
	if err != nil {
		SendBadResponse(c, "import failed with error: ", err)
		return
	}
	SendGoodResponse(c, "import succeeded", nil)
	return
}
