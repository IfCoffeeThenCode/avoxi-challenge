# avoxi-challenge
Interview coding challenge for Avoxi

# The Challenge Requirements

> The new service will be an HTTP-based API that receives an IP address and a white list of countries.  The API should return an indicator if the IP address is within the listed countries.  You can get a data set of IP addresses to country mappings from https://dev.maxmind.com/geoip/geoip2/geolite2/.

## Extra Credit

- Including a Docker file for the running service
    - Done; see Docker section below
- Including a Kubernetes YAML file for running the service in an existing cluster
    - Not yet; I'm new to Kubernetes
- Exposing the service as gRPC in addition to HTTP
    - Not yet; I'm new to gRPC
- Documenting a plan for keeping the mapping data up to date.  Extra bonus points for implementing the solution.
    - If anything, my approach is too agressive here; I'm calling
    the GeoIP API for each call (expensive both in terms of cost per call and just plain inefficient)
- Other extensions to the service you think would be worthwhile.  If you do so, please include a brief description of the feature and justification for its inclusion.  Think of this as what you would have said during the design meeting to convince your team the effort was necessary.  

# Approach

## IP address API

There is already a library built for the MaxMind API [here](https://github.com/savaki/geoip2), but in the interests of
showing off (this is an interview, after all), I will be implementing my own as if this were a [clean room implementation](https://en.wikipedia.org/wiki/Clean_room_design).

## Docker

Build service with `docker build -t avoxi .`

Run with `docker run -it --rm --name avoxi-challenge avoxi`
