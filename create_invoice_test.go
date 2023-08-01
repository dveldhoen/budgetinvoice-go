package dashweb/budgetinvoice

import "testing"

func TestCreateInvoice(t *testing.T) {
    want := "Hello, Placeholder."
    if got := CreateInvoice(); got != want {
        t.Errorf("CreateInvoice() = %q, want %q", got, want)
    }
}