provider appd {
	protocol = "https"
	host = "<controller-url>"
	port = "443"
	user = ""
	password = ""
	account = ""
}


resource appd_machineagent "ma-maac" {
	path = "<Path to MA>"
	account_access_key = ""
	sim_enabled = "true"
	ssl_enabled = "true"
}


resource "appd_healthrule" "hr-maac" {
	application_id = <application-id>
}

output "HR_details" {
	value = "${appd_healthrule.hr-maac.*.data}"
}