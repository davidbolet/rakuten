package database

import "time"

// WpPost is the type for the wordpress posts
type WpPost struct {
	ID                  uint64    `gorm:"primary_key;column:ID"`
	PostAuthor          uint64    `gorm:"column:post_author"`
	PostDate            time.Time `gorm:"column:post_date"`
	PostDateGmt         time.Time `gorm:"column:post_date_gmt"`
	PostContent         string    `gorm:"column:post_content"`
	PostTitle           string    `gorm:"not null;column:post_title"`
	PostExcerpt         string    `gorm:"not null;column:post_excerpt;default:''"`
	PostStatus          string    `gorm:"not null;column:post_status;default:'publish'"` // 'publish'
	CommentStatus       string    `gorm:"not null;column:comment_status;default:'open'"` // 'open'
	PingStatus          string    `gorm:"not null;column:ping_status;default:'open'"`    // 'open'
	PostPassword        string    `gorm:"not null;column:post_password;default:''"`
	PostName            string    `gorm:"not null;column:post_name"`
	ToPing              string    `gorm:"not null;column:to_ping;default:''"`
	Pinged              string    `gorm:"not null;column:pinged;default:''"`
	PostModified        time.Time `gorm:"not null;column:post_modified"`
	PostModifiedGmt     time.Time `gorm:"not null;column:post_modified_gmt"` //default '0000-00-00 00:00:00',
	PostContentFiltered string    `gorm:"not null;column:post_content_filtered;default:''"`
	PostParent          uint64    `gorm:"column:post_parent"`
	GUID                string    `gorm:"not null;column:guid"`
	MenuOrder           int       `gorm:"column:menu_order"`
	PostType            string    `gorm:"not null;column:post_type;default:'post'"`  //DEFAULT 'post',
	PostMimeType        string    `gorm:"not null;column:post_mime_type;default:''"` // DEFAULT '',
	CommentCount        uint64    `gorm:"column:comment_count;default:0"`            //DEFAULT 0,
}

// WpPostMeta holds metadata related to posts
type WpPostMeta struct {
	MetaID    uint64 `gorm:"primary_key;column:meta_id"`
	PostID    uint64 `gorm:"not null;column:post_id"`
	MetaKey   string `gorm:"column:meta_key"`
	MetaValue string `gorm:"column:meta_value"`
}
