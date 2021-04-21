func (client *client) request(url string) (*http.Response, error) {
    //fmt.Printf("[INFO]: %s\n", "Request to "+url)
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return &http.Response{}, fmt.Errorf("[ERR] :%s", err)
    }

    request.Header.Set("Authorization", "Bearer "+client.token)

    res, err := client.do(request)
    if err != nil {
        return res, err
    }

    return res, nil
}

func (client *client) parallelRequest(pageDetailItemCh chan pageDetailItem, url string) error {
    //fmt.Printf("[INFO]: %s\n", "Request to "+url)
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return fmt.Errorf("[ERR] :%s", err)
    }

    request.Header.Set("Authorization", "Bearer "+client.token)

    res, err := client.do(request)
    if err != nil {
        return err
    }

    var pageDetailItem pageDetailItem
    decodeBody(res, &pageDetailItem)

    pageDetailItemCh <- pageDetailItem
    return nil
}
