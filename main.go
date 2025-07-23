package main

// func main() {

// 	db := app.NewDB()
// 	validate := validator.New()

// 	categoryRepository := repository.NewCategoryRepository()
// 	categoryService := service.NewCategoryService(categoryRepository, db, validate)
// 	categoryController := controller.NewCategoryController(categoryService)

// 	router := httprouter.New()
// 	app.Router(router, categoryController)

// 	router.PanicHandler = exception.ErrorHandler

// 	server := http.Server{
// 		Addr:    "localhost:3000",
// 		Handler: middleware.NewLoggingMidleware(middleware.NewAuthMiddleware(router)),
// 	}

// 	err := server.ListenAndServe()
// 	helper.PanicIfErr(err)
// }
