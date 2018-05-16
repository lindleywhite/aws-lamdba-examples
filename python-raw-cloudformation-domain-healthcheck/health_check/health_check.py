import os
import urllib.request

def handler(event, context):
	print(event)

	domains = os.environ['DOMAINS'].split(',')
	print(domains)
	check_list = []
	for domain in domains:
		check_list.append(check_health(domain))

	return check_list

def check_health(domain):
	response = urllib.request.urlopen(domain).getcode()

	if response == 200:
		return '%s HealthCheck: %s' % (domain, response)
	else:
		return '%s HealthCheck Failed' % (domain)