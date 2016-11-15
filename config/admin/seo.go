package admin

import (
	"github.com/qor/admin"
	"github.com/qor/l10n"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/seo"
)

var SeoCollection *seo.SeoCollection

type MySeoSetting struct {
	seo.QorSeoSetting
	l10n.Locale
}

type SeoGlobalSetting struct {
	SiteName  string
	BrandName string
}

func initSeo() {
	SeoCollection = seo.New()
	SeoCollection.RegisterGlobalSetting(&SeoGlobalSetting{SiteName: "Qor Example", BrandName: "Qor"})
	SeoCollection.SettingResource = Admin.AddResource(&MySeoSetting{}, &admin.Config{Invisible: true})
	SeoCollection.RegisterSeo(&seo.Seo{
		Name: "Default Page",
	})
	SeoCollection.RegisterSeo(&seo.Seo{
		Name:     "Product",
		Settings: []string{"Name", "Code"},
		Context: func(objects ...interface{}) map[string]string {
			product := objects[0].(models.Product)
			context := make(map[string]string)
			context["Name"] = product.Name
			context["Code"] = product.Code
			return context
		},
	})
	Admin.AddResource(SeoCollection, &admin.Config{Name: "SEO Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 2})
}
