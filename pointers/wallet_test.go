package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		validateBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)
		validateBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		want := wallet.Balance()
		err := wallet.Withdraw(Bitcoin(20))
		validateBalance(t, wallet, want)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func validateBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected error but recieved none")
	}

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got error but did not expect one")
	}
}
