package acapy

// IsAlive checks if the ACA-py server is alive.
func (c *Client) IsAlive() (bool, error) {
	var result = struct {
		Alive bool `json:"alive"`
	}{}
	err := c.get("/status/live", nil, &result)
	if err != nil {
		return false, err
	}
	return result.Alive, nil
}
