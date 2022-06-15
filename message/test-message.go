package message

type Test struct {
	ID     string `json:"ID"`
	TestCd string `json:"testCd"`
}

type TestFilter struct {
	TestCd string
}

type TestResponse struct {
	ID     string
	TestCd string
}

type TestListResponse struct {
	Items []*TestResponse
}
