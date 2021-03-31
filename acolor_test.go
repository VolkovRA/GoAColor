package acolor

import (
	"fmt"
	"testing"
)

func TestApply(t *testing.T) {
	fmt.Println(Apply(Green) + "Зелёный текст" + Apply(Reset))
	fmt.Println(Apply(Blue) + "Синий текст" + Apply(Reset))
	fmt.Println(Apply(White) + "Белый текст" + Apply(Reset))
	fmt.Println(Apply(Yellow) + "Жёлтый текст" + Apply(Reset))
	fmt.Println(Apply(Red) + "Красный текст" + Apply(Reset))
}
