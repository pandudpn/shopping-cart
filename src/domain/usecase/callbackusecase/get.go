package callbackusecase

import (
	"sync"
	
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (cu *CallbackUseCase) hydrator(order *model.Order) error {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := cu.OrderProductRepo.FindProductsByOrder(order)
		if err != nil {
			errChan <- err
			return
		}

		err = cu.OrderPaymentRepo.FindPaymentByOrder(order)
		if err != nil {
			errChan <- err
			return
		}

		err = cu.OrderDeliveryRepo.FindDeliveryByOrder(order)
		if err != nil {
			errChan <- err
			return
		}

		user, err := cu.UserRepo.FindById(order.UserId)
		if err != nil {
			errChan <- err
			return
		}
		order.SetUser(user)
	}()

	select {
	case <-errChan:
		close(errChan)
		return errGetRelation
	default:
		break
	}
	wg.Wait()

	return nil
}
