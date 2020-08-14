package processor

import (
	"testing"

	"cursmedia.com/rakuten/model"
	"github.com/stretchr/testify/assert"
)

var data string = `
<result><TotalMatches>112168</TotalMatches><TotalPages>4000</TotalPages><PageNumber>1</PageNumber>
<item>
        <mid>39197</mid>
        <merchantname>Udemy</merchantname>
        <linkid>54436</linkid>
        <createdon>2016-01-15/03:54:41</createdon>
        <sku>linkshare.course.54436</sku>
        <productname>Windows 8 - What you need to know</productname>
        <category>
                <primary>Office Productivity</primary>
                <secondary>Microsoft</secondary>
        </category>
        <price currency="USD">49.99</price>
        <saleprice currency="USD">49.99</saleprice>
        <upccode/>
        <description>
                <short>Make the most of Windows 8, using both Modern UI and Desktop interfaces</short>
                <long>If you were a Windows 7, Vista or XP user, this eBook will show you everything you need to know in order to use Windows 8 efficiently. Discover Modern UI, the new</long>
        </description>
        <keywords>Microsoft Windows</keywords>
        <linkurl>http://click.linksynergy.com/link?id=2JPZ/Vwyd34&amp;offerid=507388.54436&amp;type=15&amp;murl=https%3A%2F%2Fwww.udemy.com%2Fcourse%2Fwindows-8-what-you-need-to-know%2F</linkurl>
        <imageurl>https://img-a.udemycdn.com/course/480x270/54436_bfd4_6.jpg</imageurl>
</item>
</result>
	`

func TestProcess(t *testing.T) {
	prod, err := Process(&data)
	if err != nil {
		t.Fatalf("Error should be nil but was %s", err)
	}
	assert.Equal(t, getExpectedProductResult(), prod, "Processed result should be equal to expected")
}

func getExpectedProductResult() *model.ProductResult {
	expected := model.ProductResult{}
	expected.TotalMatches = 112168
	expected.TotalPages = 4000
	expected.PageNumber = 1
	expected.Items = []model.CourseItem{}
	expectedItem := model.CourseItem{}
	expectedItem.Mid = 39197
	expectedItem.Linkid = 54436
	expectedItem.MerchantName = "Udemy"
	expectedItem.Createdon = "2016-01-15/03:54:41"
	expectedItem.Sku = "linkshare.course.54436"
	expectedItem.ProductName = "Windows 8 - What you need to know"
	expectedItem.Categories = model.TypeCategory{Primary: "Office Productivity", Secondary: "Microsoft"}
	expectedItem.Price = model.TypePrice{Price: 49.99, Currency: "USD"}
	expectedItem.SalePrice = model.TypePrice{Price: 49.99, Currency: "USD"}
	expectedItem.Description = model.TypeDescription{Short: "Make the most of Windows 8, using both Modern UI and Desktop interfaces",
		Long: "If you were a Windows 7, Vista or XP user, this eBook will show you everything you need to know in order to use Windows 8 efficiently. Discover Modern UI, the new"}
	expectedItem.Keywords = "Microsoft Windows"
	expectedItem.Linkurl = "http://click.linksynergy.com/link?id=2JPZ/Vwyd34&offerid=507388.54436&type=15&murl=https%3A%2F%2Fwww.udemy.com%2Fcourse%2Fwindows-8-what-you-need-to-know%2F"
	expectedItem.Imageurl = "https://img-a.udemycdn.com/course/480x270/54436_bfd4_6.jpg"
	expected.Items = append(expected.Items, expectedItem)

	return &expected
}
