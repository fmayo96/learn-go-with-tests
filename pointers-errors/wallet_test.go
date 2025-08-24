package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		w.Withdraw(Bitcoin(5))
		assertBalance(t, w, Bitcoin(15))
	})
	t.Run("withdraw from insufficient balance", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(15)
		assertError(t, err, withdrawError)
		assertBalance(t, w, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if err != want {
		t.Errorf("got %q want %q", err, want)
	}
}
