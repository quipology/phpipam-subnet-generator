## What's this?
This tool will provision new subnets in [phpIPAM](https://phpipam.net/).

## Requirements
- Valid app id token that is configured in phpIPAM - token must be exported with environment variable name '**PHPIPAMTOKEN**'
- YAML configuration file (for specifying desired subnets to be created) 
    YAML must include include key **CIDRs** with a list of key/value pairs which include **Name** (name/description of the subnet) and **Mask** (desired subnet mask/size). Example:
    ```
    ---
    CIDRs:
      - Name: AWS Security VPC 
        Mask: 16
      - Name: VLAN_20
        Mask: 26
      - Name: Voice VLAN
        Mask: 23
    ```
## Things to Note
You can manually set the following via flags (see usage below for details):

- YAML config file
- Master subnet Id
- Section identifier
- Base API URL

## Usage
```
Usage of ./prog
  -f string
        Specify the YAML filename. (default "cidrs.yaml")
  -m string
        Master subnet id for nested subnet. (default "231")
  -s string
        Section identifier - mandatory on add method. (default "1")
  -u string
        API base URL. (default "https://<some-domain>/api/netops")
```

## Author
[Bobby Williams](https://www.linkedin.com/in/bobby-williams-48222450)