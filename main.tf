provider appd {
	protocol = ""
	host = ""
	port = ""
	user = ""
	password = ""
	account = ""
}

variable pid {
	type = "string"
}

resource "appd_javaagent" "maac-javaagent" {
	path = ""
	jdk_path = ""
	pid = var.pid
}

resource appd_machineagent "ma-maac" {
	path = ""
	account_access_key = ""
	unique_host_id = ""
	sim_enabled = ""
	ssl_enabled = ""
}

resource "appd_healthrule" "BT-HRMaaC" {
	json_file = ""
	application_id = 
}

resource "appd_policy" "BT-PolicyMaaC" {
	json_file = ""
	application_id = 

	depends_on = [appd_healthrule.BT-HRMaaC]

}

output "HR_details" {
	value = "${appd_healthrule.BT-HRMaaC.*.data}"
}

output "Policy_details" {
	value = "${appd_policy.BT-PolicyMaaC.*.data}"
}


