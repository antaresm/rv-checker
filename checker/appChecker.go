package checker

import (
	"rv-check/models"
	"rv-check/dataProvider"
)

type AppChecker struct {
	Target models.RvCheck
}

func LoadCheckers(targets []models.RvCheck)[]AppChecker {
	appCheckers = []AppChecker{}
	for _, c := range targets {
		uCh := AppChecker{
			Target:c,
		}
		appCheckers = append(appCheckers, uCh)
	}
	return appCheckers
}

func (ch AppChecker)Check(cn chan int) {



	postStatus(ch.Target, 1)
	cn <- ch.Target.Id
}

func postStatus(ch models.RvCheck, newStatus int) {
	orm := dataProvider.OrmProvider
	ch.LastStatus = newStatus
	orm.UpdateCheck(&ch)
}