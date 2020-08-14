package database

// WpCategory models wordpress categories
type WpCategory struct {
	TermID         uint64 `gorm:"primary_key;column:term_id"`
	Name           string `gorm:"column:name"`
	Slug           string `gorm:"column:slug"`
	TermGroup      uint64 `gorm:"column:term_group"`
	TermOrder      int    `gorm:"column:term_order"`
	TermTaxonomyID uint64 `gorm:"column:term_taxonomy_id"`
	Taxonomy       string `gorm:"column:taxonomy"`
	Description    string `gorm:"column:description"`
	Parent         uint64 `gorm:"column:parent"`
	Count          uint64 `gorm:"column:count"`
}

// WpTermRelationship holds the relationships from terms and posts
type WpTermRelationship struct {
	ObjectID       uint64 `gorm:"primary_key;auto_increment:false;column:object_id"`
	TermTaxonomyID uint64 `gorm:"primary_key;auto_increment:false;column:term_taxonomy_id"`
	TermOrder      int    `gorm:"column:term_order"`
}

// CategoryRelation holds relation between linkshare and cursosrecomendados categories
type CategoryRelation struct {
	ID                int    `gorm:"primary_key;column:id"`
	LinkshareCategory string `gorm:"column:linkshare_category"`
	WpCategoryID      uint64
}
