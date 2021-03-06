# t-rex_api_wrapper
go api to wrap multi rig configuration facility handler

## Important Note
Do Not use this sofware if 8080 (or port given on cli) is open on the world. Someone with bad intention could cause damage to your rig by putting fan at 0.
Auth is planned in the futur but not yet implemented, so please use this sofware on locla network only.

## Overview

T-Rex is a versatile cryptocurrency mining software. This api wrapper allow you to manage all of those software in a signle row

## Usage

Full list of command line options:

````
-config=        is the path of the config file
````

Config File format is yaml style, there is a sample.yml as default cli command wich include all config possible:

````yaml
rig_hosts_list: # list of t-rex server adresse:port
  - "127.0.0.1:4067"
  - "127.0.0.2:4067"
  - "127.0.0.3:4067"
````

Current supported route:

    /setfan POST

    JSON : {"fan": "N"} 

       All options can be set to a comma separated list to apply different values to
       different cards. (default value for all options: 0 - not used)
        Sets GPU fan speed in percent or target temperature (auto-fan).
       Valid formats:
        {"fan": "N"}            (where N is the fan speed)
        {"fan": "t:N"}          (where N is the target core temperature)
        {"fan": "t:N[T1-T2]"}   (same as above, but with the fan speed constrained by [T1%, T2%] range)
        {"fan": "tm:N"}         (where N is the target memory temperature)
        {"fan": "tm:N[T1-T2]"}  (same as above, but with the fan speed constrained by [T1%, T2%] range)
       Example: {"fan": 45,t:67,tm:95,t:69[45-100],tm:90[50-95]
       which translates to
          GPU #0: set fan speed to 45%
          GPU #1: maintain GPU core temperature at 67C
          GPU #2: maintain GPU memory temperature at 90C
          GPU #3: maintain GPU core temperature at 69C
                  with the fan speed limited to [45%, 100%] range
          GPU #4: maintain GPU memory temperature at 90C
                  with the fan speed limited to [50%, 95%] range
       Note: fan speeds are limited to [0%, 100%] range in auto-fan mode by default.

## Json example payload

In test/json, you will find all json sample example for curl query.