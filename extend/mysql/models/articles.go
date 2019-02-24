package models

type Articles struct {
	//文章id
	Id uint `gorm:"primary_key;AUTO_INCREMENT"`
	//作者id
	Uid uint `gorm:"index"`
	//作者
	User Users `gorm:"ForeignKey:Uid;AssociationForeignKey:Id"`
	//分类ids
	CategoryIds string `gorm:"type:CHAR(255);not null;default:''"`
	//标签ids
	TagIds string `gorm:"type:CHAR(255);not null;default:''"`
	//文章标题
	Title string `gorm:"type:CHAR(50);unique_index;not null;default:''"`
	//封面图
	CoverImg string `gorm:"type:CHAR(255);not null;default:''"`
	//关键词
	KeyWord string `gorm:"type:CHAR(255);not null;default:''"`
	//文章描述
	Describe string `gorm:"type:CHAR(255);not null;default:''"`
	//文章内容
	Content string `gorm:"type:TEXT"`
	//浏览量
	PageView uint `gorm:"type:INT"`
	//评论数
	CommentsNumber uint `gorm:"type:INT"`
	//分类列表
	Category []Category
	//标签列表
	Tag []Tag
	Model
}