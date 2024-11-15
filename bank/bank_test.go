package bank_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"unit-testing/bank"
)

var _ = Describe("Bank", func() {
	var account bank.Account

	BeforeEach(func() {
		// Reset bank account before each test
		account = bank.Account{} 
	})

	Describe("Deposit", func() {
		Context("when a valid amount is deposited", func() {
			It("should increase the balance", func() {
				err := account.Deposit(100.0)
				Expect(err).To(BeNil())
				Expect(account.GetBalance()).To(Equal(100.0))
			})
		})

		Context("when an invalid amount is deposited", func() {
			It("should return an error for negative amount", func() {
				err := account.Deposit(-50.0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("deposit amount must be greater than zero"))
			})

			It("should return an error for zero amount", func() {
				err := account.Deposit(0.0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("deposit amount must be greater than zero"))
			})
		})
	})

	Describe("Withdraw", func() {
		Context("when a valid amount is withdrawn", func() {
			BeforeEach(func() {
				// Deposit before withdrawal test
				account.Deposit(250.0) 
			})

			It("should decrease the balance", func() {
				err := account.Withdraw(50.0)
				Expect(err).To(BeNil())
				Expect(account.GetBalance()).To(Equal(200.0))
			})
		})

		Context("when an invalid amount is withdrawn", func() {
			BeforeEach(func() {
				// Deposit before withdrawal test
				account.Deposit(300.0) 
			})

			It("should return an error for negative amount", func() {
				err := account.Withdraw(-50.0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("withdrawal amount must be greater than zero"))
			})

			It("should return an error for zero amount", func() {
				err := account.Withdraw(0.0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("withdrawal amount must be greater than zero"))
			})

			It("should return an error for withdraw greater than the current balance", func() {
				err := account.Withdraw(350.0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("insufficient funds"))
			})
		})
	})

	Describe("GetBalance", func() {
		It("should return the current balance", func() {
			account.Deposit(70.0)
			account.Withdraw(20.0)
			Expect(account.GetBalance()).To(Equal(50.0))
		})
	})
})
