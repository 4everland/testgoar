package testgoar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type (
	testGoArTxManager struct {
		c *http.Client

		autoMine bool
		url       string
	}

	transport struct {
		manager *testGoArTxManager
		underlyingTransport http.RoundTripper
	}
)

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Network", "arweave.testnet")
	res, err := t.underlyingTransport.RoundTrip(req)
	if req.URL.Path == "/tx" && err == nil && res.StatusCode == 200 {
		_, _ = t.manager.ResolvePool()
	}
	return res, err
}


func (m *testGoArTxManager) ReadyForMining() ([]string, error) {
	body, code, err := m.httpGet("tx/ready_for_mining")
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("get info error code: %d", code)
	}

	txs := make([]string, 0)
	err = json.Unmarshal(body, &txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}


func (m *testGoArTxManager) ResolvePool() ([]string, error)  {
	if !m.autoMine {
		return nil, nil
	}
	txIds, err := m.ReadyForMining()
	if err != nil {
		return nil, err
	}
	if len(txIds) == 0 {
		return nil, nil
	}
	err = m.Mine()
	if err != nil {
		return nil, err
	}
	return txIds, nil
}

func (m *testGoArTxManager) Mine() error {
	_, code, err := m.httpPost("mine", nil)
	if err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("get info error code: %d", code)
	}
	return nil
}


func (m *testGoArTxManager) httpGet(_path string) (body []byte, statusCode int, err error) {
	u, err := url.Parse(m.url)
	if err != nil {
		return
	}

	u.Path = path.Join(u.Path, _path)

	resp, err := m.c.Get(u.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (m *testGoArTxManager) httpPost(_path string, payload []byte) (body []byte, statusCode int, err error) {
	u, err := url.Parse(m.url)
	if err != nil {
		return
	}

	u.Path = path.Join(u.Path, _path)

	resp, err := m.c.Post(u.String(), "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	return
}



