package area

import "gorm.io/gorm"

// Versi pertama sebelum diperbaiki beserta penjelasan
/*
type (
	Area struct {
		ID        int64  `gorm:"column:id;primaryKey;"`
		AreaValue int64  `gorm:"column:area_value"`
		AreaType  string `gorm:"column:type"`
	}
)
Hanya terdapat sebuah deklarasi user defined type di sini sehingga kita tidak perlu
menggunakan declaration list seperti di atas, cukup dengan cara seperti ini:
type Area struct {
	ID        int64  `gorm:"column:id;primaryKey;"`
	AreaValue int64  `gorm:"column:area_value"`
	AreaType  string `gorm:"column:type"`
}
Parameter ketiga digunakan untuk mengindikasikan form area apa yang ingin dimasukan.
dikarenakan kita hanya memasukan satu area, kita tidak perlu menggunakan []string,
cukup dengan menggunakan string type. Selain itu, ketika menamakan sebuah variable,
hindari menggunakan nama yang sama dengan reserved keyword maupun predeclared identifier
seperti type. Argument terakhir bertipe data model.Area tidak perlu diberikan sebagai parameter
dan sebaiknya dibuat di dalam fungsi langsung. Selain itu, tipe ini juga di deklarasikan di package
yang sama sehingga tidak perlu diberi prefix model.
func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area) (err error) {
	//variable inst tidak pernah digunakan sehingga akan menyebabkan compile time error
    inst := _r.DB.Model(ar)
    var keyword ditulis dengan menggunakan lowercase
	Var area int
	deklarasi tanpa inisialisasi seperti di atas akan otomatis mengassign
	variable area ke zero value dari type yang dia miliki, yang dalam hal ini
	adalah integer. Dikarenakan zero value dari integer adalah 0, assignment
	di bawah tidak perlu dilakukan
	area = 0
    type variable diganti mengikuti penggantian pada parameter
	switch type {
        tipe data string hanya bisa dibuat dengan menggunakan quote(") atau backtick (`)
		case ‘persegi panjang’:
			deklarasi ini redundant karena kita bisa langsung melakukan kalkulasi
			dan assignment ke atribut AreaValue. Kode ini juga menyebabkan shadowing
			sehingga kita tidak bisa mengakses area variable di enclosing scope.
			Hindari menamakan variable dengan nama yang sama dengan variable di enclosing
			scope untuk menghindari ambiguitas ketika kita mengerjakan kode yang rumit
			dan panjang. param1 juga memiliki tipe data yang berbeda dengan ar.AreaValue dan
			param2 sehingga harus dikonversi menjadi in64
			var area := param1 * param2
			ar.AreaValue = area
            ilegal string: string hanya dapat dibuat dengan quote(") atau backtick(`)
			ar.AreaType = ‘persegi panjang’
			exported method di golang memiliki awalan huruf besar
			err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
        tipe data string hanya bisa dibuat dengan menggunakan quote(") atau backtick (`)
		case ‘persegi’:
			deklarasi ini redundant karena kita bisa langsung melakukan kalkulasi
			dan assignment ke atribut AreaValue. Kode ini juga menyebabkan shadowing
			sehingga kita tidak bisa mengakses area variable di enclosing scope.
			Hindari menamakan variable dengan nama yang sama dengan variable di enclosing
			scope untuk menghindari ambiguitas ketika kita mengerjakan kode yang rumit
			dan panjang. param1 juga memiliki tipe data yang berbeda dengan ar.AreaValue dan
			param2 sehingga harus dikonversi menjadi in64
			var area = param1 * param2
			ar.AreaValue = area
            ilegal string: string hanya dapat dibuat dengan quote(") atau backtick(`)
			ar.AreaType = ‘persegi’
			exported method di golang memiliki awalan huruf besar
			err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
        tidak ada variable dengan nama segitiga dan jika ingin diinterpretasikan sebagai string,
		konstanta segitiga ini juga tidak enclose dengan quote maupun backtick
		case segitiga:
			meskipun bukan deklarasi, assignment ini tetap redundant karena kita bisa
			langsung melakukan kalkulasi dan assignment ke atribut AreaValue. Selain itu, menggunakan
			variable yang ada di enclosing scope membuat kita rentan melakukan kesalahan.
			Upayakan untuk mendeklarasikan variable sedekat mungkin dengan penggunaan pertama
			variable tersebut. param1 memiliki tipe data yang berbeda dengan literal 0.5, ar.AreaValue dan
			param2 sehingga param1 harus kita konversi menjadi int64. Literal 0.5 dapat kita ganti
			dengan int64(2), dipindahkan setelah operasi perkalian dan diprepend dengan pembagian
			area = 0.5 * (param1 * param2)
			ar.AreaValue = area
            ilegal string: string hanya dapat dibuat dengan quote(") atau backtick(`)
			ar.AreaType = ‘segitiga’
			exported method di golang memiliki awalan huruf besar
			err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
		default:
			ar.AreaValue = 0
            ilegal string: string hanya dapat dibuat dengan quote(") atau backtick(`)
			ar.AreaType = ‘undefined data’
			exported method di golang memiliki awalan huruf besar
			err = _r.DB.create(&ar).Error
			if err != nil {
			r	eturn err
			}
	}
}
Dalam konteks fungsi yang dideklarasikan di atas, method call ini
membutuhkan pointer to model.Area sebagai argumen terakhir.
Selain itu, cara pembuatan string juga tidak tepat:
harus menggunakan quote atau backtick
Kita juga harus menggunakan var keyword ketika mendeklarasikan package level variable
Penempatan conditional statement seperti ini tidak dapat dilakukan di package level
err = _u.repository.InsertArea(10, 10, ‘persegi’)
if err != nil {
	log.Error().Msg(err.Error())
	err = errors.New(en.ERROR_DATABASE)
	return err
}
*/

// Versi setelah perbaikan
type Area struct {
	ID        int64  `gorm:"column:id;primaryKey;"`
	AreaValue int64  `gorm:"column:area_value"`
	AreaType  string `gorm:"column:type"`
}

type AreaRepository struct {
	DB *gorm.DB
}

func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, form string) (err error) {
	var ar Area

	switch form {
	case "persegi panjang":
		ar.AreaValue = param1 * param2
		ar.AreaType = "persegi panjang"

		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	case "persegi":
		ar.AreaValue = param1 * param2
		ar.AreaType = "persegi"

		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	case "segitiga":
		ar.AreaValue = param1 * param2 / int64(2)
		ar.AreaType = "segitiga"

		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"

		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// kode di bawah hanya valid jika tidak diletakkan di package maupun file scope
// var err = _u.repository.InsertArea(10, 10, "persegi")
// if err != nil {
// 	log.Error().Msg(err.Error())
// 	err = errors.New(en.ERROR_DATABASE)
// 	return err
// }
