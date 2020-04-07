provider appd {
	protocol = "https"
	host = ""
	port = "443"
	user = ""
	password = ""
	account = ""
}


resource appd_machineagent "ma-maac" {
	path = "<Path to MA>"
	account_access_key = ""
	unique_host_id = "MachineAgent-MaaC"
	sim_enabled = "true"
	ssl_enabled = "true"
}

resource "appd_healthrule" "hr-maac" {
	json_file = "<JSON for HR>"
	application_id = <application_id>
}

resource "appd_policy" "maac-policy" {
	json_file = "<JSON for Policy>"
	application_id = <application_id>

	depends_on = [appd_healthrule.hr-maac]
}


output "HR_details" {
	value = "${appd_healthrule.hr-maac.*.data}"
}

output "Policy_details" {
	value = "${appd_policy.maac-policy.*.data}"
}