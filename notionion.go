package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ariary/notionion/pkg/notionion"
	"github.com/jomei/notionapi"
)

func main() {
	// integration token
	token := os.Getenv("NOTION_TOKEN")
	if token == "" {
		fmt.Println("❌ Please set NOTION_TOKEN envvar with your integration token before launching notionion")
		os.Exit(92)
	}
	// page id
	pageurl := os.Getenv("NOTION_PAGE_URL")
	if pageurl == "" {
		fmt.Println("❌ Please set NOTION_PAGE_URL envvar with your page id before launching notionion (CTRL+L on desktop app)")
		os.Exit(92)
	}

	pageid := pageurl[strings.LastIndex(pageurl, "-")+1:]
	if pageid == pageurl {
		fmt.Println("❌ PAGEID was not found in NOTION_PAGEURL. Ensure the url is in the form of https://notion.so/[pagename]-[pageid]")
	}

	// Check page content
	client := notionapi.NewClient(notionapi.Token(token))

	children, err := notionion.RequestProxyPageChildren(client, pageid)
	if err != nil {
		fmt.Println("Failed retrieving page children blocks:", err)
		os.Exit(92)
	}

	active := notionion.GetProxyStatus(children)

	if active {
		fmt.Println("📶 Proxy is active")
	} else {
		fmt.Println("📴 Proxy is inactive. Activate it by checking the \"OFF\" box")
	}

	requestBlock := notionion.GetRequestBlock(children)

	if requestBlock.ID != "" {
		fmt.Println("➡️ Request block found")
	} else {
		fmt.Println("❌ Request block not found in the proxy page")
	}
	responselock := notionion.GetResponseBlock(children)
	if responselock.ID != "" {
		fmt.Println("⬅️ Response block found")
	} else {
		fmt.Println("❌ Response block not found in the proxy page")
	}

	// paragraphReq := notionion.GetRequestParagraphBlock(children)
	// if paragraphReq.ID == "" {
	// 	fmt.Println("Failed retrieving request paragraph")
	// }

	// _, err = notionion.UpdateRequestContent(client, paragraphReq.ID, "this is a test")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	codeReq := notionion.GetRequestCodeBlock(children)
	if codeReq.ID == "" {
		fmt.Println("Failed retrieving request paragraph")
	}

	_, err = notionion.UpdateRequestContent(client, codeReq.ID, "this is a test")
	if err != nil {
		fmt.Println(err)
	}

}
