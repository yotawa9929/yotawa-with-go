package helpers

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/yotawa9929/yotawa-with-go/models"
)

func ConvertContentsToMessages(contents []models.Content) (messagesToReply []linebot.Message) {
	// Set up contents for line format
	var carouselContents []models.Content
	for _, c := range contents {
		var lm linebot.Message
		if c.Category == "text" {
			lm = MakeMessageWithText(c)
		} else if c.Category == "image" {
			lm = MakeMessageWithImage(c)
		} else if c.Category == "link" {
			lm = MakeMessageWithCarousel(c)
		} else if c.Category == "sns" {
			carouselContents = append(carouselContents, c)
			// -> go to next loop b/c no linebot message to append
			continue
		} else {
			continue
		}
		messagesToReply = append(messagesToReply, lm)
	}
	//Multi-Carousel
	if len(carouselContents) > 0 {
		messagesToReply = append(messagesToReply, MakeMessageWithCarousels(carouselContents))
	}
	return messagesToReply
}

func MakeMessageWithText(c models.Content) linebot.Message {
	return linebot.NewTextMessage(c.Text)
}

func MakeMessageWithImage(c models.Content) linebot.Message {
	return linebot.NewImageMessage(c.Image, c.Image)
}

func MakeMessageWithCarousel(c models.Content) linebot.Message {
	title := c.Text
	link := c.Link
	desc := c.Link
	image := c.Image

	action := linebot.NewURITemplateAction("View", link)
	carousel := linebot.NewCarouselColumn(image, title, desc, action)
	template := linebot.NewCarouselTemplate(carousel)

	return linebot.NewTemplateMessage(title, template)
}

func MakeMessageWithCarousels(contents []models.Content) linebot.Message {
	var carousels []*linebot.CarouselColumn
	for _, c := range contents {
		title := c.Text
		link := c.Link
		desc := c.Link
		image := c.Image

		action := linebot.NewURITemplateAction("View", link)
		carousels = append(carousels, linebot.NewCarouselColumn(image, title, desc, action))
	}
	template := linebot.NewCarouselTemplate(carousels...)
	return linebot.NewTemplateMessage(contents[0].Text, template)
}
