package itswizard_module_berlin_package

import "time"

func (p *BerlinBsp) authentificate() (err error) {
	if !p.authentification {
		err = p.Authentification()
		if err != nil {
			return err
		}
		p.authentification = true
		p.createdTokenTime = time.Now()
	} else {
		if time.Now().Sub(p.createdTokenTime).Minutes() > 30 {
			err = p.Authentification()
			if err != nil {
				return err
			}
			p.createdTokenTime = time.Now()
		}
	}
	return err
}
