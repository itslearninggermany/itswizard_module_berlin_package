package itswizard_module_berlin_package

func New(clientId, clientSecret, url string) *BerlinBsp {
	p := new(BerlinBsp)
	p.grant_type = grant_type
	p.client_id = clientId
	p.client_secret = clientSecret
	p.fqdn = url
	return p
}
