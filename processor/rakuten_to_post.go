package processor

import (
	"bytes"
	"html/template"
	"log"
	"strings"
	"time"

	"cursmedia.com/rakuten/database"
	"cursmedia.com/rakuten/model"
)

// RakutenToWpConverter is the struct
type RakutenToWpConverter struct {
	Plantilla *template.Template
	AuthorID  uint64 //2
}

// NewConverter generates a converter instance
func NewConverter() *RakutenToWpConverter {
	var err error
	converter := RakutenToWpConverter{AuthorID: 2}
	converter.Plantilla, err = template.ParseFiles("resources/plantilla.html")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &converter
}

// Convert executes the conversion [^\x1F-\x7F]+
func (c *RakutenToWpConverter) Convert(course *model.CourseItem) database.WpPost {
	creation := time.Now().Add(time.Duration(-24) * time.Hour)
	result := database.WpPost{}
	result.PostAuthor = c.AuthorID
	result.PostDate = creation
	result.PostDateGmt = creation.Add(time.Duration(-1) * time.Hour)
	result.PostModified = creation
	result.PostModifiedGmt = creation.Add(time.Duration(-1) * time.Hour)
	result.PostExcerpt = course.Description.Short
	result.PostPassword = ""
	result.ToPing = "-"
	result.Pinged = "-"
	result.PostContentFiltered = "-"
	result.PostTitle = course.ProductName
	result.PostName = strings.ReplaceAll(course.Sku, ".", "-")
	var tpl bytes.Buffer
	c.Plantilla.Execute(&tpl, course)
	result.PostContent = tpl.String()
	return result
}
