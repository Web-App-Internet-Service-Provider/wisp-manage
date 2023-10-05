package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RunAPI ->route setup
func RunAPI(address string) error {

	billingStatementHandler := NewBillingStatementHandler()
	customerHandler := NewCustomerHandler()
	extensionLogHandler := NewExtensionLogHandler()
	invoiceHandler := NewInvoiceHandler()
	locationHandler := NewLocationHandler()
	messageLogsHandler := NewMessageLogsHandler()
	messageTemplatesHandler := NewMessageTemplatesHandler()
	organizationHandler := NewOrganizationHandler()
	organizationSettingHandler := NewOrganizationSettingHandler()
	paymentDetailHandler := NewPaymentDetailHandler()
	paymentMethodHandler := NewPaymentMethodHandler()
	planHandler := NewPlanHandler()
	planStationHandler := NewPlanStationHandler()
	planTypeHandler := NewPlanTypeHandler()
	poolHandler := NewPoolHandler()
	serviceHandler := NewServiceHandler()
	stationHandler := NewStationHandler()
	userHandler := NewUserHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api")

	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	billingStatementRoutes := apiRoutes.Group("/billingStatements", AuthorizeJWT())
	{
		billingStatementRoutes.GET("/add", billingStatementHandler.AddBillingStatement)
		billingStatementRoutes.DELETE("/delete", billingStatementHandler.DeleteBillingStatement)
		billingStatementRoutes.GET("/get", billingStatementHandler.GetBillingStatement)
		billingStatementRoutes.POST("/update", billingStatementHandler.UpdateBillingStatement)
		billingStatementRoutes.GET("/geteAll", billingStatementHandler.GetAllBillingStatement)
	}

	customerRoutes := apiRoutes.Group("/customers", AuthorizeJWT())
	{
		customerRoutes.POST("/update", customerHandler.UpdateCustomer)
		customerRoutes.GET("/add", customerHandler.AddCustomer)
		customerRoutes.DELETE("/delete", customerHandler.DeleteCustomer)
		customerRoutes.GET("/getAll", customerHandler.GetAllCustomer)
		customerRoutes.GET("/get", customerHandler.GetCustomer)
	}

	extensionLogRoutes := apiRoutes.Group("/extensionLog", AuthorizeJWT())
	{
		extensionLogRoutes.GET("/getAll", extensionLogHandler.GetAllExtensionLog)
		extensionLogRoutes.GET("add/", extensionLogHandler.AddExtensionLog)
		extensionLogRoutes.GET("/get", extensionLogHandler.GetExtensionLog)
		extensionLogRoutes.POST("/update", extensionLogHandler.UpdateExtensionLog)
		extensionLogRoutes.DELETE("/delete", extensionLogHandler.DeleteExtensionLog)
	}

	invoiceRoutes := apiRoutes.Group("/invoices", AuthorizeJWT())
	{
		invoiceRoutes.GET("/add", invoiceHandler.AddInvoice)
		invoiceRoutes.GET("/getAll", invoiceHandler.GetAllInvoice)
		invoiceRoutes.GET("/get", invoiceHandler.GetInvoice)
		invoiceRoutes.POST("/update", invoiceHandler.UpdateInvoice)
		invoiceRoutes.DELETE("/delete", invoiceHandler.DeleteInvoice)
	}

	locationRoutes := apiRoutes.Group("/locations", AuthorizeJWT())
	{
		locationRoutes.GET("/add", locationHandler.AddLocation)
		locationRoutes.GET("/getAll", locationHandler.GetAllLocation)
		locationRoutes.GET("/get", locationHandler.GetLocation)
		locationRoutes.POST("/update", locationHandler.UpdateLocation)
		locationRoutes.DELETE("/delete", locationHandler.DeleteLocation)
	}

	messageLogsRoutes := apiRoutes.Group("/messageLogs", AuthorizeJWT())
	{
		messageLogsRoutes.GET("/add", messageLogsHandler.AddMessageLog)
		messageLogsRoutes.GET("/get", messageLogsHandler.GetMessageLog)
		messageLogsRoutes.GET("/getAll", messageLogsHandler.GetAllMessageLog)
		messageLogsRoutes.POST("/update", messageLogsHandler.UpdateMessageLog)
		messageLogsRoutes.DELETE("/delete", messageLogsHandler.DeleteMessageLog)
	}

	messageTemplatesRoutes := apiRoutes.Group("/messageTemplates", AuthorizeJWT())
	{
		messageTemplatesRoutes.GET("/add", messageTemplatesHandler.AddMessageTemplate)
		messageTemplatesRoutes.GET("/get", messageTemplatesHandler.GetMessageTemplate)
		messageTemplatesRoutes.GET("/getAll", messageTemplatesHandler.GetAllMessageTemplate)
		messageTemplatesRoutes.POST("/update", messageTemplatesHandler.UpdateMessageTemplate)
		messageTemplatesRoutes.DELETE("delete/", messageTemplatesHandler.DeleteMessageTemplate)
	}

	organizationRoutes := apiRoutes.Group("/organization", AuthorizeJWT())
	{
		organizationRoutes.GET("/add", organizationHandler.AddOrganization)
		organizationRoutes.GET("/get", organizationHandler.GetOrganization)
		organizationRoutes.GET("/getAll", organizationHandler.GetAllOrganization)
		organizationRoutes.POST("/update", organizationHandler.UpdateOrganization)
		organizationRoutes.DELETE("/delete", organizationHandler.DeleteOrganization)
	}

	organizationSettingRoutes := apiRoutes.Group("/organizationSettings", AuthorizeJWT())
	{
		organizationSettingRoutes.GET("/add", organizationSettingHandler.AddOrganizationSetting)
		organizationSettingRoutes.GET("/get", organizationSettingHandler.GetOrganizationSetting)
		organizationSettingRoutes.GET("/getAll", organizationSettingHandler.GetAllOrganizationSetting)
		organizationSettingRoutes.POST("/update", organizationSettingHandler.UpdateOrganizationSetting)
		organizationSettingRoutes.DELETE("/delete", organizationSettingHandler.DeleteOrganizationSetting)
	}

	paymentDetailRoutes := apiRoutes.Group("/paymentDetails", AuthorizeJWT())
	{
		paymentDetailRoutes.GET("/add", paymentDetailHandler.AddPaymentDetail)
		paymentDetailRoutes.GET("/get", paymentDetailHandler.GetPaymentDetail)
		paymentDetailRoutes.GET("/getAll", paymentDetailHandler.GetAllPaymentDetail)
		paymentDetailRoutes.POST("/update", paymentDetailHandler.UpdatePaymentDetail)
		paymentDetailRoutes.DELETE("/delete", paymentDetailHandler.DeletePaymentDetail)
	}

	paymentMethodRoutes := apiRoutes.Group("/paymentMethods", AuthorizeJWT())
	{
		paymentMethodRoutes.GET("/add", paymentMethodHandler.AddPaymentMethod)
		paymentMethodRoutes.GET("/get", paymentMethodHandler.GetPaymentMethod)
		paymentMethodRoutes.GET("/getAll", paymentMethodHandler.GetAllPaymentMethod)
		paymentMethodRoutes.POST("/update", paymentMethodHandler.UpdatePaymentMethod)
		paymentMethodRoutes.DELETE("/delete", paymentMethodHandler.DeletePaymentMethod)
	}

	planRoutes := apiRoutes.Group("/plan", AuthorizeJWT())
	{
		planRoutes.GET("/add", planHandler.GetAllPlan)
		planRoutes.GET("/get", planHandler.GetPlan)
		planRoutes.GET("/getAll", planHandler.AddPlan)
		planRoutes.POST("/update", planHandler.UpdatePlan)
		planRoutes.DELETE("/delete", planHandler.DeletePlan)
	}

	planTypeRoutes := apiRoutes.Group("/planType", AuthorizeJWT())
	{
		planTypeRoutes.GET("/add", planTypeHandler.AddPlanType)
		planTypeRoutes.GET("/get", planTypeHandler.GetPlanType)
		planTypeRoutes.GET("/getAll", planTypeHandler.GetAllPlanType)
		planTypeRoutes.POST("/update", planTypeHandler.UpdatePlanType)
		planTypeRoutes.DELETE("/delete", planTypeHandler.DeletePlanType)
	}

	planStationRoutes := apiRoutes.Group("/planStations", AuthorizeJWT())
	{
		planStationRoutes.GET("/add", planStationHandler.AddPlanStation)
		planStationRoutes.GET("/get", planStationHandler.GetPlanStation)
		planStationRoutes.GET("/getAll", planStationHandler.GetAllPlanStation)
		planStationRoutes.POST("/update", planStationHandler.UpdatePlanStation)
		planStationRoutes.DELETE("/delete", planStationHandler.DeletePlanStation)
	}

	poolRoutes := apiRoutes.Group("/pools", AuthorizeJWT())
	{
		poolRoutes.GET("/add", poolHandler.AddPool)
		poolRoutes.GET("/get", poolHandler.GetPool)
		poolRoutes.GET("/getAll", poolHandler.GetAllPool)
		poolRoutes.POST("/update", poolHandler.UpdatePool)
		poolRoutes.DELETE("/delete", poolHandler.DeletePool)
	}

	serviceRoutes := apiRoutes.Group("/services", AuthorizeJWT())
	{
		serviceRoutes.GET("/add", serviceHandler.AddService)
		serviceRoutes.GET("/get", serviceHandler.GetService)
		serviceRoutes.GET("/getAll", serviceHandler.GetAllService)
		serviceRoutes.POST("/update", serviceHandler.UpdateService)
		serviceRoutes.DELETE("/delete", serviceHandler.DeleteService)
	}

	stationRoutes := apiRoutes.Group("/stations", AuthorizeJWT())
	{
		stationRoutes.GET("/add", stationHandler.AddStation)
		stationRoutes.GET("/get", stationHandler.GetStation)
		stationRoutes.GET("/getAll", stationHandler.GetAllStation)
		stationRoutes.POST("/update", stationHandler.UpdateStation)
		stationRoutes.DELETE("/delete", stationHandler.DeleteStation)
	}

	userProtectedRoutes := apiRoutes.Group("/users", AuthorizeJWT())
	{
		userProtectedRoutes.GET("/getAll", userHandler.GetAllUser)
		userProtectedRoutes.GET("/get", userHandler.GetUser)
		userProtectedRoutes.PUT("/update", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/delete", userHandler.DeleteUser)
	}

	return r.Run(address)

}
