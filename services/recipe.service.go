package services

import (
	"errors"

	"github.com/lubie-placki-be/models"
)

var me = models.User{ID: "1", Username: "krupinskij"}
var recipes = []models.Recipe{
	{
		ID:    "1",
		Title: "Murzynek",
		Image: "https://cdn.aniagotuje.com/pictures/articles/2018/03/104896-v-1000x1000.jpg",
		Time:  models.Time{Value: 180, Unit: "m"},
		Ingredients: []models.Ingredients{
			{
				Title: "",
				Ingredients: []models.Ingredient{
					{Name: "masło", Quantity: 200, Unit: "g"},
					{Name: "ekstrakt z wanilii", Quantity: 2, Unit: "łyżeczki"},
					{Name: "zmielone migdały", Quantity: 100, Unit: "g"},
				},
			},
			{
				Title: "polewa",
				Ingredients: []models.Ingredient{
					{Name: "ciemna czekolada", Quantity: 80, Unit: "g"},
				},
			},
		},
		Methods: []models.Methods{
			{
				Title: "",
				Methods: []string{
					"Ogrzać w temp. pokojowej masło, jogurt i jajka.",
					"Piekarnik nagrzać do 160 stopni C.",
					"Do misy miksera włożyć miękkie masło. Dodać cukier i ubijać mikserem przez ok. 5 - 7 minut, aż będzie jasne i puszyste.",
				},
			},
		},
		Author: me,
	},
	{
		ID:    "2",
		Title: "Piernik",
		Image: "https://wszystkiegoslodkiego.pl/storage/images/202110/piernik-weganski.jpg",
		Time:  models.Time{Value: 250, Unit: "m"},
		Ingredients: []models.Ingredients{
			{
				Title: "",
				Ingredients: []models.Ingredient{
					{Name: "masło", Quantity: 125, Unit: "g"},
					{Name: "cukier", Quantity: 250, Unit: "g"},
					{Name: "miód", Quantity: 2, Unit: "łyżki"},
				},
			},
			{
				Title: "polewa",
				Ingredients: []models.Ingredient{
					{Name: "ciemna czekolada", Quantity: 80, Unit: "g"},
				},
			},
		},
		Methods: []models.Methods{
			{
				Title: "",
				Methods: []string{
					"Piekarnik nagrzać do 175 stopni C (góra i dół bez termoobiegu).",
					"Formę keksową o wymiarach 12 x 25 cm wyłożyć papierem do pieczenia.",
					"W garnku na małym ogniu roztopić masło, dodać cukier i wymieszać. Chwilę podgrzewać aż cukier zacznie się rozpuszczać.",
				},
			},
		},
		Author: me,
	},
	{
		ID:    "3",
		Title: "Sernik",
		Image: "https://cdn.aniagotuje.com/pictures/articles/2018/11/165653-v-1000x1000.jpg",
		Time:  models.Time{Value: 210, Unit: "m"},
		Ingredients: []models.Ingredients{
			{
				Title: "ciasto kruche",
				Ingredients: []models.Ingredient{
					{Name: "mąka pszenna", Quantity: 350, Unit: "g"},
				},
			},
			{
				Title: "masa serowa",
				Ingredients: []models.Ingredient{
					{Name: "twaróg", Quantity: 1, Unit: "kg"},
				},
			},
		},
		Methods: []models.Methods{
			{
				Title: "ciasto kruche",
				Methods: []string{
					"Do mąki dodać sól, proszek do pieczenia, cukier wanilinowy, cukier oraz pokrojone w kosteczkę zimne masło.",
					"Siekać składniki na desce lub miksować mieszadłem miksera aż powstanie drobna kruszonka. Wówczas dodać jajko i żółtka i połączyć składniki w gładkie i jednolite ciasto.",
					"Uformować kulę, spłaszczyć ją i zawinąć w folię, włożyć do lodówki na ok. 30 minut.",
				},
			},
			{
				Title: "masa serowa",
				Methods: []string{
					"Połowę żółtek ubić z połową cukru pudru.",
					"Dalej ubijając dodawać po jednym żółtku i po trochu cukru pudru.",
					"Ser przepuścić 3 razy przez maszynkę razem z masłem lub dokładnie razem zmiksować ser i masło.",
				},
			},
		},
		Author: me,
	},
}

func GetRecipeById(id string) (models.Recipe, error) {
	for _, recipe := range recipes {
		if recipe.ID == id {
			return recipe, nil
		}
	}

	return models.Recipe{}, errors.New("recipe not found")
}

func GetAllRecipes() ([]models.Recipe, error) {
	return recipes, nil
}
