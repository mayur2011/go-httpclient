package gohttp

type HttpClient interface {
	Get()
	Post()
}

type httpClient struct{}

func New() HttpClient {
	client := &httpClient{}
	return client
}

func (c *httpClient) Get()  {}
func (c *httpClient) Post() {}
