import sys
import requests

API_URL = "http://ssh.jakeyee.com:9992/attempts"

def log_print(msg):
    print(msg, flush=True)

log_print("Sidecar active: Monitoring SSHD stream (Internal IP mode)...")

# Read from stdin (piped directly from sshd)
for line in sys.stdin:
    # Forward original log to Docker console
    sys.stdout.write(line)
    sys.stdout.flush()

    # Look for our specific structured log line
    if "HONEYPOT_ATTEMPT" in line:
        try:
            # Expected format: HONEYPOT_ATTEMPT|ip:1.2.3.4|user:root|pass:1234
            parts = line.split("|")
            ip = parts[1].replace("ip:", "")
            user = parts[2].replace("user:", "")
            pwd = parts[3].replace("pass:", "").strip()
            
            payload = {
                "ip": ip,
                "username": user,
                "password": pwd,
                "notes": "jaflksdjfl"
            }
            
            response = requests.post(API_URL, json=payload, timeout=5)
            log_print(f"--> API Success: {user}@{ip} (Status: {response.status_code})")
        except Exception as e:
            log_print(f"--> Parsing/API Error: {e}")