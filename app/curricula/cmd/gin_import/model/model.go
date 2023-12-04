package curricular

type Curricula struct {
	ID              uint   `gorm:"primary_key"`
	Cid             int    `gorm:"not null;comment:'课程编号'"`
	CurriculaName   string `gorm:"not null"`
	Teacher         string `gorm:"not null"`
	Type            int8   `gorm:"not null"`
	Rate            float32
	StartsNum       uint    `gorm:"column:starts_num"`
	GradeSampleSize uint    `gorm:"column:grade_sample_size"`
	TotalGrade      float32 `gorm:"column:total_grade"`
	UsualGrade      float32 `gorm:"column:usual_grade"`
	GradeR1         uint    `gorm:"column:grade_r1"`
	GradeR2         uint    `gorm:"column:grade_r2"`
	GradeR3         uint    `gorm:"column:grade_r3"`
}
