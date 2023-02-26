package main_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClientSvc(t *testing.T) {
	log.Println("==================================")
	log.Println("This will run before everything")
	log.Println("Potentially, do mocking here")
	log.Println("==================================")

	RegisterFailHandler(Fail)
	RunSpecs(t, "ClientSvc Suite")
}
