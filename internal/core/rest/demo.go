package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiPages"github.com/mohammadVatandoost/server-driven-ui/api/pages"
)

func (s *Server) GetDemoPage(c *gin.Context) {
	var widgets []*apiPages.Widget

	widgets = append(widgets, &apiPages.Widget{
		Type: apiPages.TitleType,
		Data: apiPages.TitleWidget{Title: "Demo Page"},
	})

	widgets = append(widgets, &apiPages.Widget{
		Type: apiPages.DescriptionType,
		Data: apiPages.DescriptionWidget{Description: "It is a demo page for Server Driven UI Architecture."},
	})

	widgets = append(widgets, &apiPages.Widget{
		Type: apiPages.ImageType,
		Data: apiPages.ImageWidget{Source: "https://www.shell.com/_jcr_content/par/text_over_image_caro/text_over_image_caro_1873721414/image.img.960.jpeg/1656051897866/aviation-plane.jpeg?imformat=chrome&imwidth=1280"},
	})

	widgets = append(widgets, &apiPages.Widget{
		Type: apiPages.TitleType,
		Data: apiPages.TitleWidget{Title: "List of Data"},
	})

	widgets = append(widgets, &apiPages.Widget{
		Type: apiPages.TableType,
		Data: apiPages.TableWidget{
			Titles: []string{"Column", "Column", "Column", "Column", "Column"},
			Rows: [][]string{
				{"Data", "Data", "Data", "Data", "Data"},
				{"Data", "Data", "Data", "Data", "Data"},
				{"Data", "Data", "Data", "Data", "Data"},
				{"Data", "Data", "Data", "Data", "Data"},
			},
		},
	})

	c.JSON(http.StatusOK, gin.H{
		"payload":  widgets,
	})
}
