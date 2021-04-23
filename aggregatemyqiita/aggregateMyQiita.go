package main

import (
	"fmt"
	"time"
)

func Aggregate(params Params) error {
	usersURL := "https://qiita.com/api/v2/users/" + params.UserID + "/items"

	client := client{
		token: params.Token,
	}

	res, err := client.request(usersURL)
	if err != nil {
		return err
	}

	var pageItems pageItems
	decodeBody(res, &pageItems)

	pageDetailItemCh := make(chan pageDetailItem, len(pageItems))

	go time.AfterFunc(time.Second*3, func() {
		close(pageDetailItemCh)
	})
	for _, item := range pageItems {
		go client.parallelRequest(pageDetailItemCh, "https://qiita.com/api/v2/items/"+item.ID)
	}

	for pageDetailItem := range pageDetailItemCh {
		fmt.Println("==========================================================")
		fmt.Println("ID: ", pageDetailItem.ID)
		fmt.Println("タイトル: ", pageDetailItem.Title)
		fmt.Printf("タグ: ")
		for _, tag := range pageDetailItem.Tags {
			fmt.Printf("%s, ", tag.Name)
		}
		fmt.Printf("%s", "\n")
		fmt.Println("いいね数: ", pageDetailItem.LikesCount)
		fmt.Println("閲覧数: ", pageDetailItem.PageViewsCount)
		fmt.Printf("いいね率:%.2f%%\n", pageDetailItem.likeRatio()*100)
	}
	fmt.Println("==========================================================")

	return nil
}

func main() {
	Aggregate("")
}
