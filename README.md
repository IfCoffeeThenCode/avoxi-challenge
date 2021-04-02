# avoxi-challenge
Interview coding challenge for Avoxi

# The Challenge Requirements

> The new service will be an HTTP-based API that receives an IP address and a white list of countries.  The API should return an indicator if the IP address is within the listed countries.  You can get a data set of IP addresses to country mappings from https://dev.maxmind.com/geoip/geoip2/geolite2/.

# Approach

## IP address API

There is already a library built for the MaxMind API [here](https://github.com/savaki/geoip2), but in the interests of
showing off (this is an interview, after all), I will be implementing my own as if this were a [clean room implementation](https://en.wikipedia.org/wiki/Clean_room_design).

## Docker

Build service with `docker build -t avoxi .`

Run with `docker run -it --rm --name avoxi-challenge avoxi`
