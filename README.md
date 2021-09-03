## What's this?
This tool will provision new subnets in [phpIPAM](https://phpipam.net/).

## Requirements
- Valid app id token that is configured in phpIPAM - token must be exported with environment variable name '**PHPIPAMTOKEN**'
- YAML configuration file (for specifying desired subnets to be created) 
    YAML must include include key **CIDRs** with a list of key/value pairs which include **Name** (name/description of the subnet) and **Mask** (desired subnet mask/size). Example:
    ```
    ---
    CIDRs:
      - Name: secops-vpc-prod
        Mask: 16
      - Name: voice_vlan200
        Mask: 24
      - Name: secops-vpc-sandbox
        Mask: 16
    ```
## Things to Note
You can manually set the following via flags (see usage for details):

- YAML config file
- Master subnet Id
- Base API URL

## Usage
```
Usage of ./prog
  -f string
        Specify the YAML filename. (default "cidrs.yaml")
  -m int
        Master subnet id for nested subnet. (default 231)
  -u string
        API base URL. (default "https://<some_domain>/api/netops")
```