# Nomad Usage Resource Dashboard (NURD)
NURD is a dashboard which aggregates and displays CPU and memory resource usage for each job running through specified Hashicorp Nomad servers. The dashboard also displays resources requested by each job, which can be used with resource usage to calculate waste and aid capacity planning. 

## Prerequisites
* At least one active Nomad server
* **Recommended:** A VictoriaMetrics server containing allocation level resource statistics

## Setup
1. **Configuration File**<br>
    a. **nurd/config.json**<br>
        This file contains the configuration information for the Nomad server(s) and the VictoriaMetrics server. Note, any amount of servers can be added to the `Nomad` array.

        {
            "VictoriaMetrics": {
                "URL":      URL for VictoriaMetrics server, 
                "Port":     Port for VictoriaMetrics server
            },
            "Nomad": [
                {
                    "URL":      URL for Nomad server, 
                    "Port":     Port for Nomad server
                }
            ]
        }
2. `$ git clone git@github.com:Roblox/nurd.git`
3. `$ cd nurd`
4. `$ make nurd`
5. `$ ./nurd.out`

## Usage
From `localhost:8080`, or wherever the NURD server is being hosted, the user can access several endpoints:
1. **/nurd**
2. **/nurd/jobs**
3. **/nurd/job/:job_id**