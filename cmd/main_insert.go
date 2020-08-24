package cmd

import (
	"fmt"
	"log"
	"strconv"

	"cursmedia.com/rakuten/database"
	"cursmedia.com/rakuten/model"
	"cursmedia.com/rakuten/processor"
	"cursmedia.com/rakuten/utils"
	"github.com/abadojack/whatlanggo"
)

var converter *processor.RakutenToWpConverter
var connector *database.Connector

func getMapRelations() *map[string]uint64 {
	relations := connector.ListCategoriesRelations()
	var res = make(map[string]uint64)
	for _, catRelation := range *relations {
		res[catRelation.LinkshareCategory] = catRelation.WpCategoryID
	}
	return &res
}

// Insert inserts functions into wordpress database
func Insert() {
	var language string
	connector = database.Connect()

	converter = processor.NewConverter()
	counters := make(map[string]int)
	spanishCategories := make(map[string]int)
	categoryRelation := *getMapRelations()
	for i := 1; i < 1148; i++ {
		fileN := fileName + strconv.Itoa(i) + ".xml"
		content := utils.ReadFile(fileN)
		productList, err := processor.Process(&content)
		if err != nil {
			log.Fatalf(err.Error())
		}
		emptyErr := model.ErrorCall{}
		if productList.Errors != emptyErr {
			log.Fatalf(productList.Errors.ErrorText)
		}
		for _, course := range productList.Items {
			language = checkLanguage(&course)
			counters[course.Categories.Primary] = counters[course.Categories.Primary] + 1
			if language == "Spanish" {
				spanishCategories[course.Categories.Primary] = spanishCategories[course.Categories.Primary] + 1
				spanishCategories[course.Categories.Secondary] = spanishCategories[course.Categories.Secondary] + 1
				if !connector.CheckIfExists(course.ProductName) && categoryRelation[course.Categories.Primary] > 0 {
					post := converter.Convert(&course)
					err := connector.DB.Create(&post).Error
					if err != nil {
						log.Fatalf(err.Error())
					}
					//post.GUID = "https://cursosrecomendados.com/?p=" + strconv.FormatUint(post.ID, 10)
					connector.AddRelationshipToPost(&post, categoryRelation[course.Categories.Primary])
					connector.AddRelationshipToPost(&post, 41)
					connector.AddImageToPost(&post, course.Imageurl)
					connector.AddMetaToPost(&post, "_edit_last", "2")
					connector.AddMetaToPost(&post, "platform", "Udemy")
					connector.AddMetaToPost(&post, "affiliate_link", course.Linkurl)
					connector.AddMetaToPost(&post, "regular_price", fmt.Sprintf("%.2f %s", course.Price.Price, course.Price.Currency))
				}
			}
		}
	}
	connector.Close()
	for key := range counters {

		fmt.Println(key)
	}

	//var posts []database.WpPost
	//connector.DB.Limit(10).Where("post_title like ? and post_type like 'post'", "%openoffice%").Find(&posts)
	//for _, value := range posts {
	//	fmt.Println(fmt.Sprintf("ID: %d, Title: %s", value.ID, value.PostTitle))
	//}
	//categories := connector.ListCategories()
	//for _, category := range *categories {
	//	fmt.Println(fmt.Sprintf("ID: %d, Title: %s", category.TermID, category.Name))
	//}
	/*	log.Printf("Processing Rakuten Udemy Affiliate API")
		log.Printf("1- Request token")
		rakutenClient := utils.NewRakutenClient()
		courseList := make(chan *model.CourseItem)
		log.Printf("2- Query and process products")
		for i := 1; i < 1149; i++ {

			go callAndProcess(&rakutenClient, i, courseList)
			time.Sleep(time.Second * 65)
		}
		log.Printf("Rakuten productSearch retrieved")
		saveToDatabase(courseList)*/
}

func checkLanguage(course *model.CourseItem) string {
	info := whatlanggo.Detect(course.Description.Long)
	//fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)
	return info.Lang.String()
}

func saveToDatabase(courseList chan *model.CourseItem) {
	log.Printf("Saving retrieved courses")
	count := 0
	for course := range courseList {
		log.Printf(strconv.Itoa(count) + " " + course.Description.Short)
		post := converter.Convert(course)
		connector.DB.Create(&post)
		count = count + 1
	}

}

// ProcessStringContent processes a raw xml string
func ProcessStringContent(content string) {
	var language string
	productList, err := processor.Process(&content)
	if err != nil {
		log.Fatalf(err.Error())
	}
	emptyErr := model.ErrorCall{}
	if productList.Errors != emptyErr {
		log.Fatalf(productList.Errors.ErrorText)
	}
	for _, course := range productList.Items {
		language = checkLanguage(&course)
		if language == "Spanish" {
			log.Println(course.ProductName)
			log.Println(course.Description.Long)
		}
	}
}
