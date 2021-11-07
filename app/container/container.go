package container

type Container interface {
	// BuildController akan membuat beberapa tipe untuk controller/handler dan termasuk tipe struct
	// untuk setiap panggilan, itu akan membuat instance baru.
	BuildController(code string) (interface{}, error)
	// BuildMiddleware akan membuat beberapa tipe untuk middleware dan termasuk tipe struct
	// untuk setiap panggilan, itu akan membuat instance baru
	BuildMiddleware(code string) (interface{}, error)
	// Ini hanya digunakan di dalam container dan ini adalah sub-package
	// untuk mengambil instance berdasarkan code/key container.
	Get(code string) (interface{}, bool)
	// Ini hanya digunakan di dalam container dan ini adalah sub-package
	// untuk menaruh data di dalam container berdasarkan code/key container.
	Put(code string, value interface{})
}
