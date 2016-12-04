package line

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/member/memberimp"
	"github.com/line/line-bot-sdk-go/linebot"
)

//LineApp :
type LineApp struct {
	bot           *linebot.Client
	memberService *memberimp.MemberService
}

//NewLineApp : New LineApp
func NewLineApp(channelSecret string, channelToken string, service *memberimp.MemberService) (*LineApp, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}
	return &LineApp{
		bot:           bot,
		memberService: service,
	}, nil
}

//CallbackHandler : handler
func (app *LineApp) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
			log.Println("Invalid Signature")
			log.Println("X-Line-Signature: " + r.Header.Get("X-Line-Signature"))
		} else {
			w.WriteHeader(500)
			log.Println("Unknow error")
		}
		return
	}

	log.Printf("Got events %v", events)
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				profile, err := app.bot.GetProfile(event.Source.UserID).Do()

				if message.Text == "Hi" {
					objMember := member.Member{
						LineID: profile.UserID,
						Name:   profile.DisplayName,
						Pic:    profile.PictureURL,
					}

					app.memberService.AddMember(&objMember)

					if err = app.replyText(event.ReplyToken, "ยินดีต้อนรับ "+profile.DisplayName); err != nil {
						log.Fatal(err)
					}

				} else if message.Text == "M" {
					memberObj, err := app.memberService.GetMemberByLineID(profile.UserID)
					if err != nil {
						log.Fatal(err)
					}

					id, _ := app.memberService.GetIDByLineID(memberObj.LineID)
					listMember, _ := app.memberService.AssignDay(id)

					for _, member := range listMember {
						if _, err := app.bot.PushMessage(member.LineID, linebot.NewTextMessage(memberObj.Name+" ได้เป็นสมาชิก วันเดียวกับคุณ")).Do(); err != nil {
							log.Fatal(err)
						}
					}

					memberObj, err = app.memberService.GetMemberByLineID(profile.UserID)
					if err != nil {
						log.Fatal(err)
					}

				} else {
					if err = app.replyText(event.ReplyToken, "พิมพ์ให้มันถูก ๆ หน่อย "); err != nil {
						log.Fatal(err)
					}
				}

				if err != nil {
					log.Fatal("Get Line Profile Error")
				}

				ws, err := websocket.Dial("wss://www.olipicus.com/ws", "", "https://www.olipicus.com/")
				if err != nil {
					log.Fatal(err)
				}

				wsMessage := []byte("update")
				_, err = ws.Write(wsMessage)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Send: %s\n", wsMessage)

			}
		}
	}
}

func (app *LineApp) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}
