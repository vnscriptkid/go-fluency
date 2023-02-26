package main_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	log.Println("--- Describing")
	Context("DoSomething", func() {
		log.Println("--- Context")
		It("Returns as expected", func() {
			log.Println("--- It")
			Expect(1).To(Equal(1))
		})
	})
})
