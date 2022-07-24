package main

import (
	"log"

	"github.com/gedearyarp/xendit-recon-test/interface/controller"
	"github.com/gedearyarp/xendit-recon-test/interface/repository"
	"github.com/gedearyarp/xendit-recon-test/usecase"
)

const (
	proxyFileName          = "file/transaction/proxy.csv"
	sourceFileName         = "file/transaction/source.csv"
	reconciliationFileName = "file/reconciliation/reconciliation.csv"
)

func getReconciliationController() controller.ReconciliationController {
	transactionRepo := repository.NewTransactionRepository()
	reconciliationRepo := repository.NewReconciliationRepository()
	reconciliationInteractor := usecase.NewReconciliationInteractor(reconciliationRepo, transactionRepo)
	reconciliationController := controller.NewReconciliationController(reconciliationInteractor)
	return *reconciliationController
}

func main() {
	reconciliationController := getReconciliationController()

	err := reconciliationController.ReconcilTransaction(proxyFileName, sourceFileName, reconciliationFileName)
	if err != nil {
		log.Fatal(err)
	}
}
