package mongodb

import(
	"testing"
)

func TestIsDBWork(t *testing.T){
	if session_err != nil{
		t.Error("DB EXPLOSION!!!!!!!!!")
	}
}
func TestEscape(t *testing.T){
	s := "Now 20% off, Only $88."
	Escape(&s)
	if s != "Now 20%0 off, Only %288%1" {
		t.Errorf("Bad Escape: %s || It should be: Now 20%%0 off, Only %%288%%1", s)
	}
	Unescape(&s)
	if s != "Now 20% off, Only $88." {
		t.Errorf("Bad Unescape: %s || It should be Now 20%% off, Only $88.", s)
	}
}
