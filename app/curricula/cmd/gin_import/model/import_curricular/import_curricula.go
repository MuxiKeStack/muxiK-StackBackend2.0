package import_curricular

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/dao/mysql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/model"
	"github.com/tealeg/xlsx"
	"regexp"
	"strconv"
)

func ImportCurricularByXlsx(xl *xlsx.File) error {
	for n, sheet := range xl.Sheets {
		for i, row := range sheet.Rows {
			if i > 0 {
				var curricula curricular.Curricula
				if n != 4 {
					curricula.Type = 0
				} else {
					curricula.Type = 1
				}
				for j, cell := range row.Cells {
					switch j {
					case 1:
						curricula.Cid, _ = strconv.Atoi(cell.Value)
						break
					case 2:
						curricula.CurriculaName = cell.Value
						break
					case 6:
						re := regexp.MustCompile(`^.*?/(.*?)/.*$`)
						match := re.FindStringSubmatch(cell.Value)
						if len(match) > 1 {
							curricula.Teacher = match[1]
						} else {
							curricula.Teacher = "神秘的教师"
						}
						break
					}
				}
				err := mysql.DB.Create(&curricula).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
