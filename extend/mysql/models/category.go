package models

type Category struct {
	//分类id
	Id uint `gorm:"primary_key;AUTO_INCREMENT"`
	//分类名称
	Name string `gorm:"type:CHAR(20);unique_index;not null;default:''"`
	//排序
	SortLevel uint16 `gorm:"type:TINYINT(3);not null;default:0"`
	//前端class（优先级）
	ColorClass string `gorm:"type:CHAR(20);not null;default:''"`
	//颜色代码
	ColorCode string `gorm:"type:CHAR(45);not null;default:''"`
	//热门
	IsPopular uint8 `gorm:"type:TINYINT(1);not null;default:0"`
	Model
}

type ArticlesCategory struct {
	Id uint `gorm:"primary_key;AUTO_INCREMENT"`
	//文章ID
	Aid uint `gorm:"index"`
	//实体外键约束
	Articles []Articles `gorm:"ForeignKey:Aid;AssociationForeignKey:Id"`
	//分类id
	Cid uint `gorm:"index"`
	//实体外键约束
	Category []Category `gorm:"ForeignKey:Cid;AssociationForeignKey:Id"`
	Model
}