package github.com/dveldhoen/budgetinvoice-go

import "testing"

func TestCreateInvoice(t *testing.T) {
    want := "Hello, Placeholder."
    if got := CreateInvoice(); got != want {
        t.Errorf("CreateInvoice() = %q, want %q", got, want)
    }
}