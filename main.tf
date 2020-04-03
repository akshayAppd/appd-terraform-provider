provider appd {
	protocol = "https"
	host = "ces-controller.saas.appdynamics.com"
	port = "443"
	user = "akshays"
	password = "akshays"
	account = "ces-controller"
}


resource appd_machineagent "ma-maac" {
	path = "/Users/akshasri/go/MA/"
	account_access_key = "0cff2ab5-3532-4e63-ab1a-7f02294c9f56"
	sim_enabled = "true"
	ssl_enabled = "true"
}


resource "appd_healthrule" "hr-maac" {
	application_id = 612
}

output "HR_details" {
	value = "${appd_healthrule.hr-maac.*.data}"
}