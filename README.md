# proxytld
Compares list of top-level domains (TLDs) against proxy logs (list of destination hosts). There are some TLDs that
are suspicious and this will identify if there have been any requests for sites hosted at these TLDs.

# Prerequisites
Golang 1.6+

# Usage
./proxytld -t=[path to tld file] -u=[path to url file]

# Notes
TLD file should be list of top-level domains. Example:

* .top
* .xyz
* .suspicious

URL file should be list of unique destination hosts of requests logged by proxy. Example:

* www.bing.com
* www.bloomberg.com
* api.twitter.com