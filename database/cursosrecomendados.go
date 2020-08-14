package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var instance *Connector

// Connector holds connection data
type Connector struct {
	DB              *gorm.DB
	registeredTypes []interface{}
	autoMigrate     bool
}

// Connect methods generates a new DatabaseConnection object
func Connect() *Connector {
	dbConnector := Connector{autoMigrate: true}
	pwd := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName)
	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Could not open connection to the database: %s", err)
	}
	conn.LogMode(true)
	dbConnector.DB = conn
	return &dbConnector
}

// FindPostsByTitle returns the posts from wp_posts table
func (conn *Connector) FindPostsByTitle(filterTitleWhere string) *[]WpPost {
	var result []WpPost
	conn.DB.Table("wp_posts").Where("post_title like ? and post_type like 'post'", filterTitleWhere).Scan(&result)
	return &result
}

// ListCategories returns a list of categories
func (conn *Connector) ListCategories() *[]WpCategory {
	var result []WpCategory
	conn.DB.Table("wp_terms").Joins("JOIN wp_term_taxonomy ON wp_terms.term_id = wp_term_taxonomy.term_id").Where("wp_term_taxonomy.taxonomy = 'category'").Scan(&result)
	return &result
}

// AddCategoriesToPost adds the selected category to the given post
func (conn *Connector) AddCategoriesToPost(post *WpPost, categories *[]WpCategory) {
	var relationship WpTermRelationship
	for _, category := range *categories {
		relationship = WpTermRelationship{ObjectID: post.ID, TermTaxonomyID: category.TermTaxonomyID}
		if err := conn.DB.Table("wp_term_relationships").Create(relationship).Error; err != nil {
			log.Printf("Could not create relationship %s", err.Error())
		}
	}
}

// AddRelationshipToPost adds the selected category to the given post
func (conn *Connector) AddRelationshipToPost(post *WpPost, category uint64) {

	relationship := WpTermRelationship{ObjectID: post.ID, TermTaxonomyID: category}
	if err := conn.DB.Table("wp_term_relationships").Create(relationship).Error; err != nil {
		log.Printf("Could not create relationship %s ", err.Error())
	}
}

// AddImageToPost sets the thumbnail for the course
func (conn *Connector) AddImageToPost(post *WpPost, imageURL string) {
	var imagePost WpPost
	imagePost.PostType = "attachment"
	imagePost.GUID = imageURL
	imagePost.PostMimeType = "image/jpeg"
	imagePost.PostAuthor = 2
	imagePost.PostDate = post.PostDate
	imagePost.PostStatus = "publish"

	err := conn.DB.Table("wp_posts").Create(&imagePost).Error
	if err != nil {
		log.Printf("Could not attach image to post %s ", err.Error())
	} else {
		var postMeta WpPostMeta
		postMeta.MetaKey = "_thumbnail_id"
		postMeta.MetaValue = imagePost.GUID //strconv.FormatUint(imagePost.ID, 10)
		postMeta.PostID = post.ID
		err = conn.DB.Table("wp_postmeta").Create(&postMeta).Error
		if err != nil {
			log.Printf("Could not create post meta %s ", err.Error())
		}
	}
}

// AddMetaToPost adds a meta related to the post
func (conn *Connector) AddMetaToPost(post *WpPost, metaKey string, metaValue string) {
	var postMeta WpPostMeta
	postMeta.MetaKey = metaKey
	postMeta.MetaValue = metaValue
	postMeta.PostID = post.ID
	err := conn.DB.Table("wp_postmeta").Create(&postMeta).Error
	if err != nil {
		log.Printf("Could not create post meta %s ", err.Error())
	}
}

// ListCategoriesRelations is the method that links linkshare category names to categories id
func (conn *Connector) ListCategoriesRelations() *[]CategoryRelation {
	var result []CategoryRelation
	conn.DB.Table("wp_categories_to_linkshare_cat").Scan(&result)
	return &result
}

// Close closes connection
func (conn *Connector) Close() {
	conn.DB.Close()
}
